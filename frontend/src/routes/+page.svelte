<script lang="ts">
    const apiUrl = (path: string) => `http://localhost:3000${path}`;

    const getVersion = async () => {
        const url = apiUrl('/h');
        const res = await fetch(url);
        if (!res.ok) {
            throw `Error while fetching data from ${url} (${res.status} ${res.statusText}).`;
        }
        return await res.text();
    };
</script>

<h1>Home</h1>
{#await getVersion()}
    loading...
{:then version}
    message from Server: {version}
{:catch err}
    {err}
{/await}
