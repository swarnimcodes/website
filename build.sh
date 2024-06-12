#!/bin/env bash
templ generate && go build && sudo systemctl restart gowebsite.service
