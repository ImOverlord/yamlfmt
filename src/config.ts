/**
 * yamlfmt config, based of https://github.com/google/yamlfmt/blob/main/docs/config-file.md
 */
export type Config = {
  indent?: number;
  includeDocumentStart?: boolean;
  lineEnding?: "lf" | "crlf";
  retainLineBreaks?: boolean;
  retainLineBreaksSingle?: boolean;
  disallowAnchors?: boolean;
  maxLineLength?: number;
  scanFoldedAsLiteral?: boolean;
  indentlessArrays?: boolean;
  dropMergeTag?: boolean;
  padLineComments?: number;
  trimTrailingWhitespace?: boolean;
  eofNewline?: boolean;
  stripDirectives?: boolean;
};

export const baseConfig: Config = {
  indent: 2,
  includeDocumentStart: false,
  retainLineBreaks: false,
  retainLineBreaksSingle: false,
  disallowAnchors: false,
  maxLineLength: 0,
  scanFoldedAsLiteral: false,
  indentlessArrays: false,
  dropMergeTag: false,
  padLineComments: 1,
  trimTrailingWhitespace: false,
  eofNewline: false,
  stripDirectives: false,
};

