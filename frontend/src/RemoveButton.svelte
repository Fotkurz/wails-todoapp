<script>
    const trashImg = "src/assets/images/trash.png"
</script>

<button on:click class="btn remove">
    <img src={trashImg} width="20" alt="remove todo" />
</button>


<style>
    .remove {
        background: none;
        border: none;
    }
    .remove:hover {
        background: lightblue;
        border-radius: 12px;
    }
</style>