<script lang="ts">
    import Date from "$lib/components/Date.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import * as Card from "$lib/components/ui/card";
    import {
        scriptsService,
        type FileDetails,
    } from "$lib/services/scriptsService";
    import { Play } from "lucide-svelte";
    import { onMount } from "svelte";

    let scripts: Array<FileDetails> | undefined = undefined;

    onMount(async () => {
        try {
            scripts = await scriptsService.getScripts();
        } catch (error) {
            console.error(error);
        }
    });

    async function executeScript(scriptName: string) {
        let executionResult = await scriptsService.executeScriptByName(scriptName);
        console.log('executionResult', executionResult)
    }   
</script>

<h1 class="text-2xl font-bold">Scripts</h1>

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
                            Last modified date: <Date
                                inputDate={script.modTime}
                            />
                        </span>
                        <Button
                            size="icon"
                            on:click={() => executeScript(script.name)}
                        >
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
