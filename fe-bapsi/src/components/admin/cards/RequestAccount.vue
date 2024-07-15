<template>
  <div class="cardRequest">
    <span>
      <h2>{{ data.nama }}</h2>
      <h5>{{ data.email }}</h5>
      <p>{{ data.divisi }}</p>
      <p>{{ data.region }}</p>
    </span>

    <span class="approval">
      <a href="" class="warning" @click.prevent="janganIjinkanAkses(data.id_request)">Hapus</a>
      <button class="bt-2" @click.prevent="ijinkanAkses(data.id_request)">Approve</button>
    </span>
  </div>
</template>

<script>
import {axios} from "@/axios/config.js";

export default {
  props: {
    data: Object,
  },

  methods: {
    ijinkanAkses(id_request) {
      const data = {
        request_id: id_request
      }

      axios.post("/request", data)
        .then(() => {
          window.location.reload();
        })
    },

    janganIjinkanAkses(id_request) {
      const data = {
        request_id: id_request
      }

      axios.post("/request/not", data)
        .then(() => {
          window.location.reload();
        })
    },
  }
}
</script>

<style scoped>
h5 {
  margin-bottom: 0.5rem;
}

.cardRequest {
  display: flex;
  flex-direction: column;
  max-width: 95%;
  width: 360px;
  gap: 1rem;
  box-sizing: border-box;
  padding: 1rem;

  border-radius: 8px;
  border-style: solid;
  border-color: var(--secondary);
}

span.approval {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  border-top: 1px solid;
  padding-top: 1rem;
}
</style>