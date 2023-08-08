#!/bin/sh
/migrate -path=/migrations/ -database "${DB_SOURCE}" up
