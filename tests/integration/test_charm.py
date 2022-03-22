#!/usr/bin/env python3
# Copyright 2021 Canonical Ltd.
# See LICENSE file for licensing details.
import logging
from pathlib import Path

import pytest
import yaml
from helpers import IPAddressWorkaround, is_loki_up

logger = logging.getLogger(__name__)

METADATA = yaml.safe_load(Path("./metadata.yaml").read_text())
app_name = METADATA["name"]
resources = {"loki-image": METADATA["resources"]["loki-image"]["upstream-source"]}


@pytest.mark.abort_on_fail
async def test_build_and_deploy(ops_test, loki_charm):
    """Build the charm-under-test and deploy it together with related charms.

    Assert on the unit status before any relations/configurations take place.
    """
    # build and deploy charm from local source folder
    await ops_test.model.deploy(loki_charm, resources=resources, application_name=app_name)

    async with IPAddressWorkaround(ops_test):
        await ops_test.model.wait_for_idle(apps=[app_name], status="active", timeout=1000)
        app = ops_test.model.applications[app_name]
        assert app.units[0].workload_status == "active"

    assert await is_loki_up(ops_test, app_name)
