<script lang="ts">
  import Date from "$lib/components/Date.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import * as Card from "$lib/components/ui/card";
  import { scriptsService, type FileDetails, type ScriptExecutionResponse } from "$lib/services/scriptsService";
  import { Play } from "lucide-svelte";
  import { onMount } from "svelte";
  import Terminal from "./Terminal.svelte";
  import pythonIcon from "../../assets/icons/python.svg";
  import FileIcon from "$lib/components/FileIcon.svelte";

  let scripts: Array<FileDetails> | undefined = undefined;

  let scriptsOutputs: Map<string, ScriptExecutionResponse> = new Map();

  $: scriptsExecutionResponses = Array.from(scriptsOutputs.values());

  onMount(async () => {
    try {
      scripts = await scriptsService.getScripts();
    } catch (error) {
      console.error(error);
    }
  });

  async function executeScript(scriptName: string) {
    let executionResult = await scriptsService.executeScriptByName(scriptName);
    scriptsOutputs.set(executionResult.scriptName, executionResult);
    scriptsOutputs = scriptsOutputs;
    console.log("la", scriptsExecutionResponses);
  }

  async function closeScriptTab(event: CustomEvent<{ tabName: string }>) {
    scriptsOutputs.delete(event.detail.tabName);
    console.log("close", event.detail.tabName, scriptsOutputs);
    scriptsOutputs = scriptsOutputs;
  }
</script>

<h1 class="m-4 mb-0 text-2xl font-bold">Scripts</h1>

<div class="m-4 mt-0 overflow-y-auto">
  {#if scripts}
    <ul class="flex flex-col gap-2 mt-4 overflow-y-auto">
      {#each scripts as script (script.name)}
        <li>
          <Card.Root>
            <Card.Header>
              <Card.Title class="flex items-center space-x-2">
                <FileIcon filename={script.name} class="w-6 h-6 mr-2" />
                {script.name}
              </Card.Title>
            </Card.Header>
            <Card.Content class="flex justify-between">
              <span>
                Last modified date: <Date inputDate={script.modTime} />
              </span>
              <Button size="icon" on:click={() => executeScript(script.name)}>
                <Play />
              </Button>
            </Card.Content>
          </Card.Root>
        </li>
      {/each}
    </ul>
  {:else}
    <p>Loading...</p>
  {/if}
</div>

<Terminal bind:scriptsExecutionResponses on:closeTab={closeScriptTab} />
