<script lang="ts">
    import * as Card from "$lib/components/ui/card";

    interface MetricItem {
        name: string;
        value: string | number | undefined;
    }

    let cardsData: MetricItem[] = [
        { name: "CPU Usage", value: "54 %" },
        { name: "GPU Usage", value: "13 %" },
        { name: "RAM Usage", value: "4.5 / 16 Go" },
    ];
</script>

<h1 class="text-2xl font-bold">Dashboard</h1>

<div class="grid gap-4 grid-cols-4 mt-2">
    {#each cardsData as cardData}
        <Card.Root>
            <Card.Header>
                <Card.Title>{cardData.name}</Card.Title>
            </Card.Header>
            <Card.Content>{cardData.value}</Card.Content>
        </Card.Root>
    {/each}
</div>
