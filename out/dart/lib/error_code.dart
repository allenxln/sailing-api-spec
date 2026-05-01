// GENERATED FROM codes.yaml. DO NOT EDIT.
// Edit codes.yaml then run `make gen` in sailing-api-spec.

class ErrorCode {
  ErrorCode._();
  /// 成功 (ok)
  static const int success = 0;
  /// 登录已过期 (auth status expired)
  static const int loginStatusExpired = 1001;
  /// 请求参数错误 (bad request)
  static const int badRequest = 4001;
  /// 服务器内部错误 (internal server error)
  static const int internalServerError = 5001;
  /// 词包不存在 (package not found)
  static const int packageNotFound = 20001;

  /// Returns the Chinese user-facing message for a given code.
  static String messageZh(int code) {
    switch (code) {
      case 0: return '成功';
      case 1001: return '登录已过期';
      case 4001: return '请求参数错误';
      case 5001: return '服务器内部错误';
      case 20001: return '词包不存在';
      default: return '未知错误';
    }
  }

  /// Returns the English message for a given code.
  static String messageEn(int code) {
    switch (code) {
      case 0: return 'ok';
      case 1001: return 'auth status expired';
      case 4001: return 'bad request';
      case 5001: return 'internal server error';
      case 20001: return 'package not found';
      default: return 'unknown error';
    }
  }
}
