<template>
    <div class="warp">
        <div class="cardwarp">
            <div class="card">
                <div class="card-header">{{ pokemon.name}}</div>
                <div class="card-body">
                        <img :src="pokemon.image_url"  :alt="pokemon.name"/>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props:["pokemon"],
}
</script>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    border-bottom: 0.1rem solid rgb(205, 205, 205);
    padding: 0.5rem;
    margin-right: 1rem;
    margin-left: 1rem;
    font-size: 1.5rem;
}
.card-body{
    display: flex;
    align-items: center;
    padding: 1rem;

}

.cardwarp{
    display: flex;
    justify-content: space-evenly;
    align-items: center;
}
.card {
    display: flex;
    flex-direction: column;
    background-color: rgb(255, 255, 255);
    justify-content: center;
    border-radius: 0.625rem;
    height: 100%;
    width: 20rem;
    border: 4px solid rgba(0, 69, 138, 0.5);
    box-shadow: rgb(0 0 0 / 55%)
}

.warp {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-height: 70%;
    overflow: auto;
    padding: 1rem;
}


@media only screen and (max-width: 600px) {
    .cardwarp{
        width: 100%;
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        overflow: auto;
        margin-top: 2rem;
        padding-bottom: 2.5rem;
    }
    .card{
        display: flex;
        flex-direction: column;
        background-color: rgb(255, 255, 255);
        justify-content: center;
        max-width: 55rem;
        width: 50%;
        border-radius: 0.625rem;
        min-height: 11rem;
        overflow: hidden;
        border: 0.1rem solid rgb(205, 205, 205);
        box-shadow: rgb(0 0 0 / 15%) 0px 2px 10px 0px;
    }
}

</style>
