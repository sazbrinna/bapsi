<template>
  <h1 class="title gap">Tambah Divisi</h1>

  <section>
    <div class="wrapper">
      <p>Divisi hanya untuk asisten</p>
      <form @submit.prevent="submitData">
        <span>
          <label for="divisi">Divisi</label>
          <input type="text" id="divisi" placeholder="Divisi..." v-model="divisi" required>
        </span>

        <input type="submit" class="bt-1">
      </form>
    </div>
  </section>
</template>

<script>
import { axios } from "@/axios/config.js";

export default {
  data() {
    return {
      divisi: "",
      role: "asisten",
    }
  },
  methods: {
    submitData() {
      const data = {
        divisi: this.divisi,
        group: this.role,
      };

      axios.post("/divisi", data)
        .then(() => {
          window.scrollTo({
            top: 0,
            behavior: 'smooth'
          });

          this.$router.push({ name: 'KonfigurasiView' });
        })
    }
  }
}
</script>