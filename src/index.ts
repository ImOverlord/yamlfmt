import { readFile } from "node:fs/promises";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import "./wasm_exec"; // This adds global.Go
import { baseConfig } from "./config";
import type { Config } from "./config";

type Result<T> = {
  Error: string;
  Value: T
}

type FormatFunction = (buffer: string, config: Config) => Result<string>;
declare global {
  class go {
    constructor();
    run(instance: any): void;
    importObject: any;
  }
  var Go: typeof go;
  var format: FormatFunction;
}


let formatYamlFunc: FormatFunction | null = null;
async function setup(): Promise<FormatFunction> {
  // Load the WebAssembly wasm file
  const CURRENT_DIR = dirname(fileURLToPath(import.meta.url));
  const wasmBuffer = await readFile(resolve(CURRENT_DIR, "./main.wasm"));
  // Set up the WebAssembly module instance
  const go = new global.Go();
  const { instance } = await WebAssembly.instantiate(wasmBuffer, go.importObject);
  go.run(instance);
  return global.format;
}

/**
 * formatYAML
 * @description Format YAML string
 * @param buffer YAML content string
 * @param config {Config} Taken from yamlfmt config
 * @throws {Error} if yamlfmt returns an error or fail to start go wasm
 */
export async function formatYAML(buffer: string, config?: Config): Promise<string> {
  if (!formatYamlFunc) {
    formatYamlFunc = await setup();
  }
  config = Object.assign(baseConfig, config);
  const result = formatYamlFunc(buffer, config);
  if (!result.Error) {
    return result.Value;
  }
  throw new Error(result.Error);
}
