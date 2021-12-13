#!/bin/sh

# Copy service file, incase if there are any changes
cp test-cicd.service /etc/systemd/system/test-cicd.service
# reload configurations incase if service file has changed
systemctl daemon-reload
# restart the service
systemctl restart test-cicd