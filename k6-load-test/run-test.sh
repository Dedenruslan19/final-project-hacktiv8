#!/bin/bash
# Load environment variables from .env file

# Check if .env file exists
if [ ! -f .env ]; then
    echo ".env file not found!"
    echo "Copy .env.example to .env and fill in your values:"
    echo "  cp .env.example .env"
    exit 1
fi

# Load .env file
export $(cat .env | grep -v '^#' | xargs)

echo "Environment variables loaded:"
echo "   BASE_URL: $BASE_URL"
echo "   SESSION_ID: $SESSION_ID"
echo "   ITEM_ID: $ITEM_ID"
echo "   AUTH_TOKEN: ${AUTH_TOKEN:0:20}..."

# Run k6 test
k6 run \
  -e BASE_URL="$BASE_URL" \
  -e SESSION_ID="$SESSION_ID" \
  -e ITEM_ID="$ITEM_ID" \
  -e AUTH_TOKEN="$AUTH_TOKEN" \
  "$@"
