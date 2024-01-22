<script lang="ts">
  import Date from "$lib/components/Date.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import * as Card from "$lib/components/ui/card";
  import { scriptsService, type FileDetails, type ScriptExecutionResponse } from "$lib/services/scriptsService";
  import { Cross, Play, X } from "lucide-svelte";
  import { onMount } from "svelte";
  import { quintOut } from "svelte/easing";
  import { fly, slide } from "svelte/transition";

  let scripts: Array<FileDetails> | undefined = undefined;

  let scriptsOutputs: Map<string, ScriptExecutionResponse> = new Map();

  onMount(async () => {
    try {
      scripts = await scriptsService.getScripts();
    } catch (error) {
      console.error(error);
    }
  });

  async function executeScript(scriptName: string) {
    let executionResult = await scriptsService.executeScriptByName(scriptName);
    console.log("executionResult", executionResult);
    scriptsOutputs.set(executionResult.scriptName, executionResult);
    scriptsOutputs = scriptsOutputs;
    console.log("size", scriptsOutputs.size);
  }

  async function closeScriptOutput(scriptName: string) {
    scriptsOutputs.delete(scriptName);
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
              <Card.Title>{script.name}</Card.Title>
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

{#if scriptsOutputs.size > 0}
  <div transition:fly={{ delay: 250, duration: 300, y: 500, opacity: 0.5, easing: quintOut }} class="flex-grow flex flex-col h-1/2 border-t">
    {#each [...scriptsOutputs] as [scriptName, scriptResponse]}
      <div class="flex justify-between">
        <h2>{scriptName}</h2>
        <Button variant="ghost" size="icon" on:click={() => closeScriptOutput(scriptName)}>
          <X />
        </Button>
      </div>
      <!-- TODO: terminal component -->
      <div class="whitespace-pre overflow-y-auto flex-grow bg-black text-white font-mono">
        {@html scriptResponse.output.split("\n\r").join("<br>")}
      </div>
    {/each}
  </div>
{/if}