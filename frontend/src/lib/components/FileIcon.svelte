<script lang="ts">
  import pythonIcon from "../../assets/icons/python.svg";
  import powershellIcon from "../../assets/icons/powershell.svg";

  let className: string | undefined = undefined;
  export { className as class };

  export let filename: string;

  const iconDictionary: { [extension: string]: string } = {
    py: pythonIcon,
    ps1: powershellIcon,
  };

  let iconPath = (() => {
    const fileExtension = filename.split(".").pop() as string;
    return iconDictionary[fileExtension] || "";
  })();
</script>

<!-- TODO: unknown file extension -->
<img src={iconPath} alt="Python Icon" class={className} />
