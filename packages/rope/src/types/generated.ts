// This file is auto-generated from the OpenAPI schema. Do not edit manually.

/**
 * Query parameters for schema requests
 */
export interface SchemaQueryParams {
  chainId?: number;
  protocol?: string;
  action?: string;
  from?: string;
  search?: IntentSearchQueryParam[];
}

export interface ActionsOption {
  icon?: ActionsOptionIcon;
  info?: ActionsOptionInfo;
  label?: string;
  name?: string;
  value?: string;
}

export interface ActionsOptionIcon {
  default?: string;
  secondary?: string;
}

export interface ActionsOptionInfo {
  label?: string;
  value?: string;
}

export interface ActionsOptions {
  complex?: { [key: string]: ActionsOption[] };
  simple?: ActionsOption[];
}

export interface ActionsProtocolMetadata {
  chains?: ReferencesNetwork[];
  icon?: string;
  tags?: string[];
}

export interface ActionsProtocolSchema {
  metadata?: ActionsProtocolMetadata;
  schema?: { [key: string]: ActionsSchema };
}

export interface ActionsSchema {
  coils?: { [key: string]: string };
  options?: { [key: string]: ActionsOptions };
  sentence?: string;
  type?: string;
}

export interface ApiKeyApiKeyCreateRequest {
  /** Rate limit for this API key */
  rateLimit?: number;
  /** Role for this API key (user or admin) */
  role?: string;
  /** ID of the socket to associate with this key */
  socketId?: string;
}

export interface CoilSlice {
  index?: number;
  length?: string;
  name?: string;
  start?: string;
  type?: string;
  typeId?: number;
}

export interface CoilUpdate {
  slice?: CoilSlice;
  start?: string;
}

export interface HealthHealthResponse {
  /** The health status of the API */
  status?: string;
}

export interface IntentSearchQueryParam {
  /** Index of the search parameter */
  index?: number;
  /** Value of the search parameter */
  value?: string;
}

export interface KillKillResponse {
  /** Whether the kill switch is active */
  killed?: boolean;
}

export interface ModelsApiKey {
  id?: string;
  key?: string;
  rateLimit?: number;
  role?: string;
  socketId?: string;
}

export interface ModelsIntent {
  accessList?: TypesAccessList;
  chainId?: number;
  createdAt?: string;
  endAt?: string;
  frequency?: number;
  from?: string;
  gasLimit?: number;
  id?: string;
  inputs?: TypesInputs;
  locked?: boolean;
  nextSimulationAt?: string;
  options?: TypesOptions;
  periodEndAt?: string;
  runs?: ModelsRun[];
  saved?: boolean;
  startAt?: string;
  status?: string;
  value?: string;
}

export interface ModelsRun {
  data?: ModelsRunOutputData;
  error?: string;
  errors?: string[];
  from?: string;
  gasEstimate?: number;
  id?: string;
  intentId?: string;
  livePlugs?: SignatureLivePlugs;
  livePlugsId?: string;
  status?: string;
  to?: string;
  value?: string;
}

export interface ModelsRunOutputData {
  decoded?: any;
  raw?: string;
}

export interface ReferencesNetwork {
  chainIds?: number[];
  icon?: { default?: string };
  name?: string;
}

export type SaveIntentListResponse = ModelsIntent[];

export interface SignatureLivePlugs {
  chainId?: number;
  createdAt?: string;
  data?: string;
  deletedAt?: string;
  from?: string;
  id?: string;
  intentId?: string;
  plugs?: SignaturePlugs;
  signature?: string;
  updatedAt?: string;
}

export interface SignaturePlug {
  data?: string;
  meta?: any;
  selector?: number;
  to?: string;
  updates?: CoilUpdate[];
  value?: string;
}

export interface SignaturePlugs {
  plugs?: SignaturePlug[];
  salt?: string;
  socket?: string;
  solver?: string;
}

export type TypesAccessList = TypesAccessTuple[];

export interface TypesAccessTuple {
  address?: string;
  storageKeys?: string[];
}

export type TypesInputs = { [key: string]: any }[];

export interface TypesOptions {
  [key: string]: any;
}

