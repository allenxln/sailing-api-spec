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
  /** 邮箱已被注册 (email already registered) */
  EmailAlreadyRegistered: 21003,
  /** 邮箱或密码错误 (email or password wrong) */
  LoginFailed: 21004,
  /** 验证码发送过于频繁，请稍后再试 (verification code sent too frequently) */
  VerifyCodeSendLimited: 21005,
  /** 尝试次数过多，请稍后再试 (too many login attempts) */
  TooManyLoginAttempts: 21006,
  /** 树节点操作非法（移到自身子树下等） (tree node operation invalid) */
  TreeNodeInvalid: 21007,
  /** 同级下已存在同名节点 (sibling node name conflict) */
  TreeNodeNameConflict: 21008,
  /** 题目不存在 (question not found) */
  QuestionNotFound: 21009,
  /** 验证链接无效或已过期 (verification token invalid or expired) */
  VerifyTokenInvalid: 21010,
  /** 知识树不存在 (tree not found) */
  TreeNotFound: 21011,
  /** 同维度知识树已存在 (tree with same dimensions already exists) */
  TreeAlreadyExists: 21012,
  /** 该知识点已挂载在此节点 (knowledge point already mounted on this node) */
  KnowledgeAlreadyMounted: 21013,
  /** 试卷不存在 (paper not found) */
  PaperNotFound: 21014,
  /** 题目导入批次不存在 (question import not found) */
  QuestionImportNotFound: 21015,
  /** 题目草稿不存在 (question draft not found) */
  QuestionDraftNotFound: 21016,
  /** 题目解析失败 (question parse failed) */
  ParseFailed: 21017,
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
    case 21003: return '邮箱已被注册';
    case 21004: return '邮箱或密码错误';
    case 21005: return '验证码发送过于频繁，请稍后再试';
    case 21006: return '尝试次数过多，请稍后再试';
    case 21007: return '树节点操作非法（移到自身子树下等）';
    case 21008: return '同级下已存在同名节点';
    case 21009: return '题目不存在';
    case 21010: return '验证链接无效或已过期';
    case 21011: return '知识树不存在';
    case 21012: return '同维度知识树已存在';
    case 21013: return '该知识点已挂载在此节点';
    case 21014: return '试卷不存在';
    case 21015: return '题目导入批次不存在';
    case 21016: return '题目草稿不存在';
    case 21017: return '题目解析失败';
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
    case 21003: return 'email already registered';
    case 21004: return 'email or password wrong';
    case 21005: return 'verification code sent too frequently';
    case 21006: return 'too many login attempts';
    case 21007: return 'tree node operation invalid';
    case 21008: return 'sibling node name conflict';
    case 21009: return 'question not found';
    case 21010: return 'verification token invalid or expired';
    case 21011: return 'tree not found';
    case 21012: return 'tree with same dimensions already exists';
    case 21013: return 'knowledge point already mounted on this node';
    case 21014: return 'paper not found';
    case 21015: return 'question import not found';
    case 21016: return 'question draft not found';
    case 21017: return 'question parse failed';
    default: return 'unknown error';
  }
}
