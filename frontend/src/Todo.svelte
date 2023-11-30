<script lang="ts">
    export let text = '';
    export let index: number;

    function removeMe(i: number): number {
        return i
    }
</script>

<li class="todo">
    {text}
    <button class="btn" on:click={() => removeMe(index)}>x</button>
</li>

<style>
    .todo {
        display: flex;
    }

</style>