<script lang="ts">
    export let inputDate: Date;
    //TODO: cleaner component

    function formatDate(aDate: Date) {
        const date = new Date(aDate);
        const year = date.getFullYear();
        const month = ("0" + (date.getMonth() + 1)).slice(-2); // Add 1 because months are 0-indexed
        const day = ("0" + date.getDate()).slice(-2);
        const hours = ("0" + date.getHours()).slice(-2);
        const minutes = ("0" + date.getMinutes()).slice(-2);
        const seconds = ("0" + date.getSeconds()).slice(-2);

        return `${year}/${month}/${day} ${hours}:${minutes}:${seconds}`;
    }
</script>

{formatDate(inputDate)}
