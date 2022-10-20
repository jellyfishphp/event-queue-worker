#!/bin/bash
json=$(cat << EOF
{
    "event1": [
        {
            "identifier": "identifier1",
            "type": "async"
        }
    ],
    "event2": [
        {
            "identifier": "identifier2",
            "type": "async"
        }
    ]
}
EOF
)

echo $json
