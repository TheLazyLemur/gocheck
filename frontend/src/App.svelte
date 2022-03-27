<script>
    import { onMount } from "svelte";

    let cpuUsage = 0;
    let memoryUsage = 0;

    onMount(() => {
        wails.Events.On("cpu", (data) => {
            cpuUsage = data.avg;
        });

        wails.Events.On("mem", (data) => {
            memoryUsage = data.avg;
        });
    });
</script>

<main>
    <div class="App">
        <header class="App-header">
            <div class="h-screen w-screen">
                <div>
                    <label for="cpu">Cpu Usage: {cpuUsage}%</label>
                    <progress id="cpu" value={cpuUsage} max="100">
                        {cpuUsage}%
                    </progress>
                </div>
                <div>
                    <label for="mem">Mem Usage {memoryUsage}%</label>
                    <progress id="mem" value={memoryUsage} max="100">
                        {memoryUsage}%
                    </progress>
                </div>
            </div>
        </header>
    </div>
</main>
