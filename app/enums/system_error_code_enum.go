package enums

type SystemErrorCode string

const ERROR_CODE_1 SystemErrorCode = "ERROR_CODE_1"
const GENERAL_CLIENT_ERROR SystemErrorCode = "BAD_REQUEST"
const GENERAL_SERVER_ERROR SystemErrorCode = "INTERNAL_SERVER_ERROR"
const NOT_FOUND_ERROR SystemErrorCode = "NOT_FOUND"
const PERSONAL_ACCESS_CLIENT_NOT_FOUND = "PERSONAL_ACCESS_CLIENT_NOT_FOUND"
const UNAUTHORIZED_ERROR SystemErrorCode = "Unauthorized"
