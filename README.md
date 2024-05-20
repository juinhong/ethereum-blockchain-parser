# ethereum-blockchain-parser

## get_address_transactions
- **Usage**: Get a list of inbound or outbound transactions for an address
- **URL**: `/get_address_transactions`
- **Method**: `GET`
- **Auth required**: No
- **Permissions required**: None
### Request
- **address** (required | string): The address to get the list of inbound or outbound transactions
- **from_block** (optional | int): The block number from which the logs (transactions) should start being collected
- **to_block** (optional | int): The block number at which the log (transaction) collection should stop

### Response 
- **transactions** (repeated): The list of inbound or outbound transactions for an address
- **debugMessage** (string): The error message returned to the client if the API call fails

## observe_address_transactions
- **Usage**: Add address to observer
- **URL**: `/observe_address_transactions`
- **Method**: `GET`
- **Auth required**: No
- **Permissions required**: None
### Request
- **address** (required | string): The address to be added to observer

### Response
- **observeSuccess** (bool): A bool to tell if the address has been added to observer
- **debugMessage** (string): The error message returned to the client if the API call fails

## get_current_block
- **Usage**: Get last parsed block
- **URL**: `/get_current_block`
- **Method**: `GET`
- **Auth required**: No
- **Permissions required**: None
### Request
No parameters required.

### Response
- **blockNumber** (int): Last parsed block number
- **debugMessage** (string): The error message returned to the client if the API call fails