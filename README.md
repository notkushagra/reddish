# reddish

## Overview
This is part of my reinventing the wheel self-series where I create my own worse versions of already existing technologies.

## Motivation
Reason: because I'm a developer who **loves to understand how things work under the hood**
> I'll reinvent my own wheel here even though its slower, more expensive and probably more square than round.

## Project Goals
Core idea—store key-value pairs in memory, handle basic commands (GET, SET, DEL), and optionally write to disk.

### Things to keep in mind
0. Focus on building MVP first.
1. Probably use go routines to handle multiple instances.
2. Try to make it concurrent.

### Implementation Steps
1. Use RESP protocol
2. Build your own parser.
3. Be redis-cli client compatible.
4. Build a simple web app for demonstration.

## Prerequisites
1. Install redis on your own machine so you have access to redis-cli.
