<script lang="ts">
    export let text = '';

</script>

<li class="todo">
    {text}
</li>

<style>
    .todo {
        word-wrap: break-word;
        color: black;
        border-radius: 5px;
        width: 100%;
        min-height: 15px;
        text-align: left;
        padding: 3px 3px 3px 5px;
    }

</style>