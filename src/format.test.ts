import { describe, expect, it } from "vitest";
import { formatYAML } from "./index.ts"

const yaml = `
app:
  versions:
    - name: cake
      version: 1.0.0
    - name: cookie
      version: 2.0.0
`

const expectYaml = `---
app:
  versions:
  - name: cake
    version: 1.0.0
  - name: cookie
    version: 2.0.0
`

describe("Format", () => {

  it("should take in account config", async () => {
    const output = await formatYAML(yaml, { includeDocumentStart: true, indentlessArrays: true});
    expect(output).toEqual(expectYaml)
  })

});