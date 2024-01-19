<script lang="ts">
  import Router from "svelte-spa-router";
  import svelteLogo from "./assets/svelte.svg";
  import viteLogo from "./assets/vite.svg";
  import Counter from "./lib/Counter.svelte";
  import A from "./routes/A.svelte";
  import B from "./routes/B.svelte";
  import Button from "$lib/components/ui/button/button.svelte";

  const routes = {
    "/a": A,
    "/b": B,
  };
</script>

Root

<Button href="#/">Go root</Button>
<Button href="#/a">Go to A</Button>
<Button href="#/b">Go to A</Button>

<Router {routes} />
