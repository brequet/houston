<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import { Cog, LayoutDashboard, Terminal } from "lucide-svelte";
  import type { ComponentType } from "svelte";
  import Router, { location } from "svelte-spa-router";
  import favicon from "./assets/svelte.svg";
  import DashboardPage from "./routes/DashboardPage.svelte";
  import ScriptsPage from "./routes/scripts-page/ScriptsPage.svelte";
  import ServicesPage from "./routes/services-page/ServicesPage.svelte";

  const navItems: {
    name: string;
    path: string;
    component: ComponentType;
    icon: ComponentType;
  }[] = [
    {
      name: "Dashboard",
      path: "/dashboard",
      component: DashboardPage,
      icon: LayoutDashboard,
    },
    {
      name: "Scripts",
      path: "/scripts",
      component: ScriptsPage,
      icon: Terminal,
    },
    {
      name: "Services",
      path: "/services",
      component: ServicesPage,
      icon: Cog,
    },
  ];

  const routes = navItems.reduce(
    (accumulator, navItem) => {
      accumulator[navItem.path] = navItem.component;
      return accumulator;
    },
    {} as Record<string, ComponentType>,
  );
</script>

<svelte:head>
  <link rel="icon" type="image/svg" href={favicon} />
</svelte:head>

<div class="w-screen h-screen flex flex-row">
  <nav class="w-60 border-r py-2 px-2">
    <h2 class="text-xl font-semibold py-2 px-4 tracking-tight">Overview</h2>
    <div class="grid items-start gap-2">
      {#each navItems as navItem}
        <Button href="#{navItem.path}" class="flex-grow justify-start" variant={$location === navItem.path ? "secondary" : "ghost"}>
          <svelte:component this={navItem.icon} class="mr-2 w-4 h-4" />
          {navItem.name}
        </Button>
      {/each}
    </div>
  </nav>

  <main class="flex-grow shrink overflow-auto flex flex-col">
    <Router {routes} />
  </main>
</div>
