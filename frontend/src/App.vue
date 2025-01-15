<template>
  <div>
    <h1>Web Scraper Application Objective</h1>
    <section v-if="count===0">
      <h3>pls type url</h3>
      <input type="text" v-model="url"  />
      <div>
        <button class="submit" @click="submit">send</button>
      </div>


    </section>
    <section v-else class="container">
      <div v-if="pokemonsC.length===0">
        loading ....
      </div>
      <div v-else>
      <div>Count : {{count}}</div>

        <section class="offset">


        <nav aria-label="Page navigation example">
          <ul class="pagination">
            <li class="page-item">
              <button type="button" class="page-link" v-if="page != 1" @click="page--"> Previous </button>
            </li>
            <li class="page-item">
              <button type="button" class="page-link" v-for="pageNumber in pages.slice(page-1, page+5)" @click="page = pageNumber"> {{pageNumber}} </button>
            </li>
            <li class="page-item">
              <button type="button"  @click="page++" v-if="page < pages.length" class="page-link"> Next </button>
            </li>
          </ul>
        </nav>
        </section>

        <div class="" v-for="pokemon in pokemonsC" :key="pokemon.name">
        <Card :pokemon="pokemon"></Card>
      </div>

      </div>
    </section>

  </div>

</template>

<script setup>

import { ref,onMounted,computed ,watch} from 'vue'
import axios  from "axios";
import Card from "@/components/Card.vue";


const url = ref("https://pokemondb.net/pokedex/all")
const pokemons = ref([])
const count = ref(null)

const page = ref(1)
const perPage = ref(9)
const pages = ref([])

const pokemonsC = computed(() => paginate(pokemons.value))

watch(pokemons, (newPokemons) => {
  setPages()
})
const setPages =() =>{
  let numberOfPages = Math.ceil(pokemons.value.length / perPage.value);
  for (let index = 1; index <= numberOfPages; index++) {
    pages.value.push(index);
  }
}

const paginate = (pokemons)=> {
  let from = (page.value * perPage.value) - perPage.value;
  let to = (page.value * perPage.value);
  return  pokemons.slice(from, to);
}

const getPokemons=()=>{
  axios.get(import.meta.env.VITE_URL || "http://127.0.0.1:7777"+"/pokemons")
    .then(res=>{
      pokemons.value = res.data.data.pokemons
      count.value  = res.data.data.count
    })
    .catch(err =>console.log(err))
}

onMounted(() => {
 getPokemons()
  setPages()
})

const submit = () =>{
  if (url.value ==='')return
  axios.post(import.meta.env.VITE_URL || "http://127.0.0.1:7777"+"/scrape", {url:url.value})
    .then(res=>{
      if (res.data.status ===200)getPokemons()
    })
    .catch(err =>console.log(err))
}


</script>

<style scoped>
input{
  background-color: rgb(255, 255, 255);
  justify-content: center;
  border-radius: 0.625rem;
  height: 2.8rem;
  width: 20rem;
  border: 4px solid rgba(0, 69, 138, 0.5);
}

button.submit{
  background-color: rgb(255, 255, 255);
  justify-content: center;
  border-radius: 0.625rem;
  height: 2.8rem;
  width: 20rem;
  font-size: medium;
  border: 4px solid rgba(0, 69, 138, 0.5);
  margin-top: 2rem;
  cursor: pointer;
}

ul {
  width: 100%;
  padding-left: 0;
}

li {
  list-style: none;
  display: inline-block;
  width: calc(100% / 3);
  height: 20px;
  text-align: center;
}

button.page-link {
  display: inline-block;
}
button.page-link {
  font-size: 20px;
  color: #29b3ed;
  font-weight: 500;
}
.offset{
  width: 700px !important;
  margin: 20px auto;
}

</style>
