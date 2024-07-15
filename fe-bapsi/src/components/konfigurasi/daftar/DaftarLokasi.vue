<template>
  <section>
    <div class="wrapper">
      <h2 class="title">Daftar Loket</h2>
      <p class="link" @click="PushRoute('TambahLoket')">Tambah loket</p>

      <div class="table">
        <table>
          <thead>
            <tr>
              <th>No</th>
              <th>Loket</th>
              <th>Latitude</th>
              <th>Longitude</th>
              <th>Aksi</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="(region, index) in regions" :key="index" :class="{ gray: index % 2 == 0 }">
              <td>{{ index + 1 }}</td>
              <td>{{ region.region }}</td>
              <td>{{ region.latitude }}</td>
              <td>{{ region.longitude }}</td>
              <td>
                <span class="green" @click="editRegion(region)">edit</span> |
                <span class="red" @click="confirmDeleteRegion(region.id_region, region.region)">delete</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>

  <section v-if="editRegions.status" class="absolute">
    <!-- sementara edit lokasi langsung di sini -->

    <form @submit.prevent="submitLoket">
      <h2>Edit Loket</h2>

      <span>
        <label for="loket">Loket</label>
        <input type="text" id="loket" v-model="editRegions.region" />
      </span>

      <span>
        <label for="latitude">Latitude</label>
        <input type="text" id="latitude" v-model="editRegions.latitude" />
      </span>

      <span>
        <label for="longitude">Longitude</label>
        <input type="text" id="longitude" v-model="editRegions.longitude" />
      </span>

      <input type="submit" class="bt-2">
    </form>
  </section>

  <ConfirmationCard v-if="konfirmasiHapus" :nama="namaRegion" @kembali="handlerKembali()" @hapus="DeleteRegion()"/>
  <AlertCard v-if="message" :msg="message" />
</template>

<script>
import AlertCard from "@/components/alert/AlertCard.vue";
import ConfirmationCard from "@/components/alert/ConfirmationalCard.vue";
import { axios } from "@/axios/config.js";

export default {
  data() {
    return {
      regions: [],
      message: "",

      editRegions: {
        status: false,
        id_region: "",
        region: "",
        latitude: "",
        longitude: "",
      },

      konfirmasiHapus: false,
      idRegion: "",
      namaRegion: "",
    }
  },

  components: {
    ConfirmationCard,
    AlertCard,
  },

  mounted() {
    axios.get("/location/all")
      .then((response) => {
        this.regions = response.data;
      })
      .catch(() => {
          this.message = "Gagal menerima lokasi";
          setTimeout(() => {
            this.message = "";
          }, 2000)
        });
  },

  methods: {
    editRegion(regionData) {
      this.editRegions.status = true;
      this.editRegions.id_region = regionData.id_region;
      this.editRegions.region = regionData.region;
      this.editRegions.latitude = regionData.latitude;
      this.editRegions.longitude = regionData.longitude;
    },

    PushRoute(pageName) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ name: pageName })
    },

    submitLoket() {
      const data = {
        id_loket: this.editRegions.id_region,
        loket: this.editRegions.region,
        latitude: this.editRegions.latitude.toString(),
        longitude: this.editRegions.longitude.toString(),
      }

      axios.post("/loket/edit", data)
        .then(() => {
          window.location.reload();
        })
        .catch(() => {
          this.message = "Gagal mengubah data loket";
          setTimeout(() => {
            this.message = "";
          }, 2000)
        });
    },

    confirmDeleteRegion(idRegion, namaRegion) {
      this.konfirmasiHapus = true;
      this.idRegion = idRegion;
      this.namaRegion = namaRegion;
    },

    handlerKembali() {
      this.konfirmasiHapus = false;
      this.idRegion = "";
      this.namaRegion = "";
    },

    DeleteRegion() {
      axios.delete(`/region/${this.idRegion}`)
      .then(() => {
        this.handlerKembali();
      })
      .catch(() => {
          this.message = "Gagal menghapus data loket";
          setTimeout(() => {
            this.message = "";
          }, 2000)
        });
    }
  }
}
</script>

<style scoped>
.absolute {
  position: fixed;
  background-color: var(--white);
  width: 100%;
  height: 100vh;
  top: 0;
}
</style>

<style scoped>
td span {
  cursor: pointer;
}
</style>