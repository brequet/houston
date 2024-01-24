<script lang="ts">
  import Button from "$lib/components/ui/button/button.svelte";
  import { type ScriptExecutionResponse } from "$lib/services/scriptsService";
  import { X } from "lucide-svelte";
  import { createEventDispatcher } from "svelte";
  import { quintOut } from "svelte/easing";
  import { fly } from "svelte/transition";

  export let scriptsExecutionResponses: ScriptExecutionResponse[];

  const dispatch = createEventDispatcher();

  $: console.log("CHANGED", scriptsExecutionResponses);

  let selectedScriptTabIndex: number | undefined = undefined;

  $: if (scriptsExecutionResponses.length > 0 && selectedScriptTabIndex === undefined) {
    selectedScriptTabIndex = getDefaultScriptTabIndex();
    console.log("CHANGED here ?", scriptsExecutionResponses);
  }

  function getDefaultScriptTabIndex(): number | undefined {
    if (scriptsExecutionResponses.length === 0) {
      return undefined;
    }
    return scriptsExecutionResponses.length - 1;
  }

  function selectTab(tabIndex: number) {
    selectedScriptTabIndex = tabIndex;
  }

  async function closeScriptTab(tabIndex: number) {
    dispatch("closeTab", { tabName: scriptsExecutionResponses[tabIndex].scriptName });
    if (selectedScriptTabIndex === tabIndex) {
      selectedScriptTabIndex = undefined;
    }
  }
</script>

{#if scriptsExecutionResponses.length > 0 && selectedScriptTabIndex !== undefined}
  <div transition:fly={{ duration: 300, y: 500, opacity: 0.5, easing: quintOut }} class="flex-grow flex flex-col h-1/2 border-t">
    <div class="bg-slate-50 flex flex-row items-center">
      {#each scriptsExecutionResponses as scriptExecutionResponse, tabIndex}
        {@const tabScriptName = scriptExecutionResponse.scriptName}
        <button class="flex flex-row items-center pl-1 {tabIndex === selectedScriptTabIndex ? 'bg-slate-300' : 'bg-slate-200'}" on:click={() => selectTab(tabIndex)}>
          <h2>{tabScriptName}</h2>
          <Button variant="ghost" size="icon" on:click={() => closeScriptTab(tabIndex)}>
            <X />
          </Button>
        </button>
      {/each}
    </div>

    <!-- TODO: terminal component -->
    <div class="whitespace-pre overflow-y-auto flex-grow bg-black text-white font-mono">
      {@html scriptsExecutionResponses[selectedScriptTabIndex]?.output.split("\n\r").join("<br>")}
    </div>
  </div>
{/if}
