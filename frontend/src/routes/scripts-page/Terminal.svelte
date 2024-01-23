<script lang="ts">
  import Button from "$lib/components/ui/button/button.svelte";
  import { type ScriptExecutionResponse } from "$lib/services/scriptsService";
  import { X } from "lucide-svelte";
  import { createEventDispatcher } from "svelte";
  import { quintOut } from "svelte/easing";
  import { fly } from "svelte/transition";

  export let scriptsExecutionResponses: ScriptExecutionResponse[];

  const dispatch = createEventDispatcher();

  let selectedScriptTab: ScriptExecutionResponse | undefined = undefined;

  $: if (scriptsExecutionResponses.length > 0 && selectedScriptTab === undefined) {
    selectedScriptTab = getDefaultScriptTab();
  }

  function getDefaultScriptTab(): ScriptExecutionResponse | undefined {
    if (scriptsExecutionResponses.length === 0) {
      return undefined;
    }
    return scriptsExecutionResponses[0];
  }

  function selectTab(scriptName: string) {
    selectedScriptTab = scriptsExecutionResponses.find((e) => e.scriptName === scriptName);
  }

  async function closeScriptTab(scriptName: string) {
    dispatch("closeTab", { tabName: scriptName });
    if (selectedScriptTab?.scriptName === scriptName) {
      selectedScriptTab = undefined;
    }
  }
</script>

{#if scriptsExecutionResponses.length > 0 && selectedScriptTab !== undefined}
  <div transition:fly={{ delay: 250, duration: 300, y: 500, opacity: 0.5, easing: quintOut }} class="flex-grow flex flex-col h-1/2 border-t">
    <div class="bg-slate-50 flex flex-row items-center">
      {#each scriptsExecutionResponses as scriptExecutionResponse}
        {@const tabScriptName = scriptExecutionResponse.scriptName}
        <button class="flex flex-row items-center pl-1 {tabScriptName === selectedScriptTab.scriptName ? 'bg-slate-300' : 'bg-slate-200'}" on:click={() => selectTab(tabScriptName)}>
          <h2>{tabScriptName}</h2>
          <Button variant="ghost" size="icon" on:click={() => closeScriptTab(tabScriptName)}>
            <X />
          </Button>
        </button>
      {/each}
    </div>

    <!-- TODO: terminal component -->
    <div class="whitespace-pre overflow-y-auto flex-grow bg-black text-white font-mono">
      {@html selectedScriptTab.output.split("\n\r").join("<br>")}
    </div>
  </div>
{/if}
