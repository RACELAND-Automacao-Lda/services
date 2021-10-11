import * as fx from "m3o/function";

// Deploy a group of functions
async function DeployAfunction() {
  let functionService = new fx.FunctionService(process.env.MICRO_API_TOKEN);
  let rsp = await functionService.deploy({
    entrypoint: "helloworld",
    name: "my-first-func",
    project: "tests",
    repo: "github.com/m3o/nodejs-function-example",
  });
  console.log(rsp);
}

await DeployAfunction();
