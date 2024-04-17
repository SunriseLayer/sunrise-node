#!/bin/bash

set -e

if [ "$1" = 'sunrise' ]; then
    echo "Initializing Sunrise Node with command:"

    if [[ -n "$NODE_STORE" ]]; then
        echo "sunrise "${NODE_TYPE}" init --p2p.network "${P2P_NETWORK}" --node.store "${NODE_STORE}""
        sunrise "${NODE_TYPE}" init --p2p.network "${P2P_NETWORK}" --node.store "${NODE_STORE}"
    else
        echo "sunrise "${NODE_TYPE}" init --p2p.network "${P2P_NETWORK}""
        sunrise "${NODE_TYPE}" init --p2p.network "${P2P_NETWORK}"
    fi

    echo ""
    echo ""
fi

echo "Starting Sunrise Node with command:"
echo "$@"
echo ""

exec "$@"
