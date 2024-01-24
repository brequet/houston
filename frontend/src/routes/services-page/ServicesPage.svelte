<script lang="ts">
  import * as Card from "$lib/components/ui/card";
  import { servicesService, type Service } from "$lib/services/servicesService";
  import { onMount } from "svelte";

  let services: Array<Service> | undefined = undefined;

  onMount(async () => {
    try {
      services = await servicesService.getServices();
    } catch (error) {
      console.error(error);
    }
  });
</script>

<h1 class="m-4 mb-0 text-2xl font-bold">Services</h1>

<div class="m-4 mt-0 overflow-y-auto">
  {#if services}
    <ul class="flex flex-col gap-2 mt-4 overflow-y-auto">
      {#each services as service (service.name)}
        <li>
          <Card.Root>
            <Card.Header>
              <Card.Title class="flex items-center space-x-2">
                {service.name}
              </Card.Title>
            </Card.Header>
            <Card.Content class="flex justify-between">
              <ul>
                <li>{service.description}</li>
                <li>{service.status}</li>
                <li>{service.processId}</li>
              </ul>
            </Card.Content>
          </Card.Root>
        </li>
      {/each}
    </ul>
  {:else}
    <p>Loading...</p>
  {/if}
</div>
