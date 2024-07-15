<template>
  <TitlePart title="Invitation" description="available" />

  <section>
    <div class="wrapper">
      <form @submit.prevent="submitData">
        <span>
          <label for="email">Email</label>
          <input id="email" type="text" v-model="email" placeholder="Email...">
        </span>

        <span>
          <label for="role">Role</label>
          <select id="role" v-model="role" @change="getDivisi()">
            <option value="">Pilih role...</option>
            <option value="admin">Admin</option>
            <option value="staff">Staff</option>
            <option value="asisten">Asisten</option>
          </select>
        </span>

        <span v-if="role == 'asisten'">
          <label for="divisi">Divisi</label>
          <select id="divisi" v-model="divisi">
            <option value="">Pilih divisi...</option>
            <option v-for="(divisi, index) in divisis" :key="index" :value="divisi.divisi"> {{ divisi.divisi }}
            </option>
          </select>
        </span>

        <span v-if="role == 'asisten'">
          <label for="region">Lokasi Loket</label>
          <select id="region" v-model="region">
            <option value="">Pilih region...</option>
            <option v-for="(region, index) in regions" :key="index" :value="region.region">{{ region.region }}</option>
          </select>
        </span>

        <input type="submit" class="bt-1">
      </form>
    </div>
  </section>

</template>

<script>
import { axios } from "@/axios/config.js";
import TitlePart from "@/components/TitlePart.vue";

export default {
  components: {
    TitlePart,
  },

  data() {
    return {
      role: "",
      email: "",

      divisis: [],
      regions: [],

      divisi: "",
      region: "",
    }
  },

  created() {
    this.getRegions();
  },

  methods: {
    getDivisi() {
      if (this.role != 'asisten') {
        return
      }
      
      axios.get("/divisi")
        .then((response) => {
          this.divisis = response.data;
        })
    },

    getRegions() {
      axios.get("/location/all")
        .then((response) => {
          this.regions = response.data;
        })
    },

    submitData() {
      const data = {
        email: this.email,
        division: this.division,
        region: this.region,
        role: this.role
      };

      axios.post("/admin/invitation", data)
        .then(() => {
        })
    }
  }
}
</script>