<template>
  <main class="flex flex-col justify-center items-center gap-8 py-12 px-8">

    <h1 class="text-3xl text-blue-500">Welcome to Home page</h1>

    <!-- Display Person's name or text -->
    <p>Form API : {{ data?.name || "Click the Button 👇" }}</p>
    <!-- Click the button to (re)fetch the data from the API -->
    <button class="p-2 bg-lime-600 text-whiet font-bold rounded-md" @click="fetchData">{{ data ? "Refresh" : "Get data" }}</button>

    <div class="flex flex-col gap-4">
      <!-- v-model binds the input to the reference "name". It is both a Getter and a Setter -->
      <input class="p-2 rounded-md" placeholder="no data.." type="text" v-model="name">
      <!-- Fire the Post Request to update the name -->
      <button class="p-2 bg-lime-600 text-whiet font-bold rounded-md" @click="update">Update Name</button>
    </div>
    <p>Name model : {{ name }}</p>
  </main>
</template>
<script lang="ts" setup>
import { ref } from 'vue';

type Person = {
  name: string
  age: number
  email: string
}

type PostPersonName = {
  name: string
}

const data = ref<Person | null>(null)

const name = defineModel<string>("fetching..")

async function fetchData() {
  const prom = await fetch("http://localhost:8888/person")
  const res: Awaited<Person> = await prom.json()
  data.value = res
  name.value = res.name
}

async function update() {
  const data: PostPersonName = {
    name: name.value!
  }

  const resp = await fetch("http://localhost:8888/person", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
    },
    body: JSON.stringify(data)
  })

  console.log(resp)
}

</script>