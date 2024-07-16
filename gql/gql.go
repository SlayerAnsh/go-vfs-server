package gql

const GQL_URL = "https://api.andromedaprotocol.io/graphql/testnet"

const CHAIN_QUERY = `query CHAIN_QUERY($chain:String!){chainConfigs{config(identifier:$chain){kernelAddress}}}`

const SMART_QUERY = `query SMART_CONTRACT($contract: String!, $query: String!) {ADO {adoSmart(address: $contract, query: $query) {queryResult}}}`

const KERNEL_KEY_ADDRESS_QUERY = `{\"key_address\":{\"key\":\"%s\"}}`

const VFS_RESOLVE_PATH_QUERY = `{\"resolve_path\":{\"path\":\"%s\"}}`

const ADO_TYPE_QUERY = `{\"type\":{}}`

const PRIMITIVE_KEY_QUERY = `{\"get_value\":{\"key\":\"%s\"}}`
