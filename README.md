<div align="center">
  yamlfmt NodeJS SDK
  <br />
  <br />
  <a href="https://github.com/ImOverlord/yamlfmt/issues/new?assignees=&labels=bug&template=01_BUG_REPORT.md&title=bug%3A+">Report a Bug</a>
  ·
  <a href="https://github.com/ImOverlord/yamlfmt/issues/new?assignees=&labels=enhancement&template=02_FEATURE_REQUEST.md&title=feat%3A+">Request a Feature</a>
  .
  <a href="https://github.com/ImOverlord/yamlfmt/issues/new?assignees=&labels=question&template=04_SUPPORT_QUESTION.md&title=support%3A+">Ask a Question</a>
</div>

<div align="center">
<br />

[![Project license](https://img.shields.io/github/license/ImOverlord/yamlfmt.svg?style=flat-square)](LICENSE)

[![Pull Requests welcome](https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=flat-square)](https://github.com/ImOverlord/yamlfmt/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
[![code with love by ImOverlord](https://img.shields.io/badge/%3C%2F%3E%20with%20%E2%99%A5%20by-ImOverlord-ff1414.svg?style=flat-square)](https://github.com/ImOverlord)

</div>



---

## About

When building tools that manipulate YAML files, I would encounter issues where they wouldn't respect the formatting rules. Instead of building another YAML formatter, I wanted to try to include the formatter directly from JS. Using WASM, this tool calls [yamlfmt](https://github.com/google/yamlfmt) directly with the specified config.

## Getting Started

### Prerequisites

> Explain what is needed to run the project

### Installation

```shell
npm i @imoverlord/yamlfmt
```

## Usage

```typescript
import { formatYAML } from "./index.ts"

const yaml = `
app:
  versions:
    - name: cake
      version: 1.0.0
    - name: cookie
      version: 2.0.0
`

await formatYAML(yaml)
```

`formatYAML` will start the Go wasm if needed and the directly pass the string to yamlfmt go formatter.

You can specify a config if needed. The config should support all options from [yamlfmt config](https://github.com/google/yamlfmt/blob/main/docs/config-file.md).

```typescript
import { formatYAML } from "./index.ts"

const yaml = `
app:
  versions:
    - name: cake
      version: 1.0.0
    - name: cookie
      version: 2.0.0
`

await formatYAML(yaml, { indentlessArrays: true })
```

### Config

| Name                        | Type           | Default | Description |
|:----------------------------|:---------------|:--------|:------------|
| `indent`                    | int            | 2       | The indentation level in spaces to use for the formatted yaml. |
| `includeDocumentStart`    | bool           | false   | Include `---` at document start. |
| `lineEnding`               | `lf` or `crlf` | `crlf` on Windows, `lf` otherwise | Parse and write the file with "lf" or "crlf" line endings. This setting will be overwritten by the global `line_ending`. |
| `retainLineBreaks`        | bool           | false   | Retain line breaks in formatted yaml. |
| `retainLineBreaksSingle` | bool           | false   | (NOTE: Takes precedence over `retain_line_breaks`) Retain line breaks in formatted yaml, but only keep a single line in groups of many blank lines. |
| `disallowAnchors`          | bool           | false   | If true, reject any YAML anchors or aliases found in the document. |
| `maxLineLength`           | int            | 0       | Set the maximum line length ([see note below](#max_line_length)). if not set, defaults to 0 which means no limit. |
| `scanFoldedAsLiteral`    | bool           | false   | Option that will preserve newlines in folded block scalars (blocks that start with `>`). |
| `indentlessArrays`         | bool           | false   | Render `-` array items (block sequence items) without an increased indent. |
| `dropMergeTag`            | bool           | false   | Assume that any well formed merge using just a `<<` token will be a merge, and drop the `!!merge` tag from the formatted result. |
| `padLineComments`         | int            | 1       | The number of padding spaces to insert before line comments. |
| `trimTrailingWhitespace`  | bool           | false   | Trim trailing whitespace from lines. |
| `eofNewline`               | bool           | false   | Always add a newline at end of file. Useful in the scenario where `retainLineBreaks` is disabled but the trailing newline is still needed. |
| `stripDirectives`          | bool           | false   | [YAML Directives](https://yaml.org/spec/1.2.2/#3234-directives) are not supported by this formatter. This feature will attempt to strip the directives before formatting and put them back. [Use this feature at your own risk.](#strip_directives) |

Table was taken directly from [yamlfmt config documentation](https://github.com/google/yamlfmt/blob/main/docs/config-file.md). The keys were adapted to be typescript friendly camel case instead of pascal case.

## License

This project is licensed under the **Apache-2.0 license**.

See [LICENSE](LICENSE) for more information.
