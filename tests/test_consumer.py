# Copyright 2020 Canonical Ltd.
# See LICENSE file for licensing details.

import json
import unittest
from pathlib import Path
from unittest.mock import PropertyMock, patch

from charms.loki_k8s.v0.loki import AlertRuleError, LokiConsumer
from ops.charm import CharmBase
from ops.framework import StoredState
from ops.testing import Harness

LABELED_ALERT_RULES = [
    {
        "name": "loki_20ce8299-3634-4bef-8bd8-5ace6c8816b4_promtail-k8s_alerts",
        "rules": [
            {
                "alert": "HighPercentageError",
                "expr": 'sum(rate({juju_model="loki", juju_model_uuid="20ce8299-3634-4bef-8bd8-5ace6c8816b4", juju_application="promtail-k8s"} |= "error" [5m])) by (job)\n  /\nsum(rate({app="foo", env="production"}[5m])) by (job)\n  > 0.05\n',
                "for": "10m",
                "labels": {
                    "severity": "page",
                    "juju_model": "loki",
                    "juju_model_uuid": "20ce8299-3634-4bef-8bd8-5ace6c8816b4",
                    "juju_application": "promtail-k8s",
                },
                "annotations": {"summary": "High request latency"},
            }
        ],
    }
]

ONE_RULE = {
    "alert": "HighPercentageError",
    "expr": 'sum(rate({%%juju_topology%%} |= "error" [5m])) by (job)\n  /\nsum(rate({app="foo", env="production"}[5m])) by (job)\n  > 0.05\n',
    "for": "10m",
    "labels": {"severity": "page"},
    "annotations": {"summary": "High request latency"},
}


class FakeConsumerCharm(CharmBase):
    _stored = StoredState()

    def __init__(self, *args, **kwargs):
        super().__init__(*args)
        self._port = 3100
        self.loki_consumer = LokiConsumer(self, "logging")

    @property
    def _loki_push_api(self) -> str:
        loki_push_api = f"http://{self.unit_ip}:{self.charm._port}/loki/api/v1/push"
        data = {"loki_push_api": loki_push_api}
        return json.dumps(data)

    @property
    def unit_ip(self) -> str:
        """Returns unit's IP"""
        return "10.1.2.3"


class TestLokiConsumer(unittest.TestCase):
    def setUp(self):
        self.harness = Harness(FakeConsumerCharm)
        self.addCleanup(self.harness.cleanup)
        self.harness.set_leader(True)
        self.harness.begin()

    def test__on_logging_relation_changed_no_leader(self):
        self.harness.set_leader(False)
        rel_id = self.harness.add_relation("logging", "promtail")
        self.harness.add_relation_unit(rel_id, "promtail/0")
        self.assertEqual(self.harness.update_relation_data(rel_id, "promtail/0", {}), None)

    def test__on_logging_relation_changed_no_unit(self):
        rel_id = self.harness.add_relation("logging", "promtail")
        self.assertEqual(self.harness.update_relation_data(rel_id, "promtail", {}), None)

    @patch("charms.loki_k8s.v0.loki.LokiConsumer._labeled_alert_groups", new_callable=PropertyMock)
    def test__on_logging_relation_changed(self, mock_alert_rules):
        mock_alert_rules.return_value = LABELED_ALERT_RULES
        self.harness.set_leader(True)
        rel_id = self.harness.add_relation("logging", "promtail")
        self.harness.add_relation_unit(rel_id, "promtail/0")
        self.harness.update_relation_data(
            rel_id,
            "promtail/0",
            {"data": '{"loki_push_api": "http://10.1.2.3:3100/loki/api/v1/push"}'},
        )

    def test__label_alert_expression(self):
        labeled_alert = self.harness.charm.loki_consumer._label_alert_expression(ONE_RULE.copy())
        self.assertTrue("juju_model" in labeled_alert["expr"])
        self.assertTrue("juju_model_uuid" in labeled_alert["expr"])
        self.assertTrue("juju_application" in labeled_alert["expr"])

    def test__label_alert_topology(self):
        labeled_alert_topology = self.harness.charm.loki_consumer._label_alert_topology(
            ONE_RULE.copy()
        )
        self.assertTrue("juju_model" in labeled_alert_topology["labels"])
        self.assertTrue("juju_model_uuid" in labeled_alert_topology["labels"])
        self.assertTrue("juju_application" in labeled_alert_topology["labels"])

    def test__validate_alert_rule(self):
        thefile = Path("rulefile.rule")
        self.assertIsNone(
            self.harness.charm.loki_consumer._validate_alert_rule(ONE_RULE.copy(), thefile)
        )

        with self.assertRaises(AlertRuleError):
            rule_1 = ONE_RULE.copy()
            del rule_1["alert"]
            self.harness.charm.loki_consumer._validate_alert_rule(rule_1, thefile)

        with self.assertRaises(AlertRuleError):
            rule_2 = ONE_RULE.copy()
            del rule_2["expr"]
            self.harness.charm.loki_consumer._validate_alert_rule(rule_2, thefile)

        with self.assertRaises(AlertRuleError):
            rule_3 = ONE_RULE.copy()
            rule_3["expr"] = "Missing Juju topology placeholder"
            self.harness.charm.loki_consumer._validate_alert_rule(rule_3, thefile)
