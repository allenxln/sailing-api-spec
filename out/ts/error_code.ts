// GENERATED FROM codes.yaml. DO NOT EDIT.
// Edit codes.yaml then run `make gen` in sailing-api-spec.

export const ErrorCode = {
  /** 成功 (ok) */
  Success: 0,
  /** 登录已过期 (auth status expired) */
  LoginStatusExpired: 1001,
  /** 权限不足 (permission denied) */
  PermissionDenied: 1003,
  /** 请求参数错误 (bad request) */
  BadRequest: 4001,
  /** 服务器内部错误 (internal server error) */
  InternalServerError: 5001,
  /** 词包不存在 (package not found) */
  PackageNotFound: 20001,
  /** 知识点名称已存在 (knowledge name already exists) */
  KnowledgeNameConflict: 21001,
  /** 验证码错误或已过期 (verification code invalid or expired) */
  VerifyCodeInvalid: 21002,
} as const;

/** 中文用户可见文案。 */
export function messageZh(code: number): string {
  switch (code) {
    case 0: return '成功';
    case 1001: return '登录已过期';
    case 1003: return '权限不足';
    case 4001: return '请求参数错误';
    case 5001: return '服务器内部错误';
    case 20001: return '词包不存在';
    case 21001: return '知识点名称已存在';
    case 21002: return '验证码错误或已过期';
    default: return '未知错误';
  }
}

/** English message. */
export function messageEn(code: number): string {
  switch (code) {
    case 0: return 'ok';
    case 1001: return 'auth status expired';
    case 1003: return 'permission denied';
    case 4001: return 'bad request';
    case 5001: return 'internal server error';
    case 20001: return 'package not found';
    case 21001: return 'knowledge name already exists';
    case 21002: return 'verification code invalid or expired';
    default: return 'unknown error';
  }
}
