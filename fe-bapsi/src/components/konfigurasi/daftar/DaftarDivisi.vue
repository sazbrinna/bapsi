<template>
  <section>
    <div class="wrapper">
      <h2 class="title">Daftar Divisi</h2>
      <p class="link" @click="PushRoute('TambahDivisi')">Tambah divisi</p>

      <div class="table">
        <table>
          <thead>
            <tr>
              <th>No</th>
              <th>Divisi</th>
              <th>Aksi</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="(divisi, index) in divisis" :key="index" :class="{ gray: index % 2 == 0 }">
              <td>{{ index + 1 }}</td>
              <td>{{ divisi.divisi }}</td>
              <td>
                <span class="green" @click="BukaEditDivisi(divisi)">edit</span> | <span class="red"
                  @click="confirmDeleteDivisi(divisi.id_divisi, divisi.divisi)">delete</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>

  <section v-if="formEdit" class="absolute">
    <form @submit.prevent="submitDivisi">
      <h2>Edit Divisi</h2>

      <span>
        <label for="divisi">Divisi</label>
        <input type="text" id="divisi" v-model="editDivisi.divisi" />
      </span>
      <input type="submit" class="bt-2" value="update">
    </form>
  </section>

  <AlertCard v-if="message" :msg="message" />
  <ConfirmationCard v-if="konfirmasiHapus" :nama="namaDivisi" @kembali="handlerKembali()" @hapus="DeleteDivisi()"/>
</template>

<script>
import { axios } from "@/axios/config.js";
import AlertCard from "@/components/alert/AlertCard.vue";
import ConfirmationCard from "@/components/alert/ConfirmationalCard.vue";

export default {
  data() {
    return {
      divisis: [],
      message: "",
      editDivisi: [],

      formEdit: false,
      konfirmasiHapus: false,
      iddivisi: "",
      namaDivisi: "",
    }
  },

  components: {
    AlertCard,
    ConfirmationCard,
  },

  mounted() {
    axios.get("/divisi")
      .then((response) => {
        this.divisis = response.data.sort((a, b) => a.divisi.localeCompare(b.divisi));
      })
  },

  methods: {
    PushRoute(pageName) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ name: pageName })
    },

    confirmDeleteDivisi(idDivisi, namaDivisi) {
      this.konfirmasiHapus = true;
      this.iddivisi = idDivisi;
      this.namaDivisi = namaDivisi;
    },

    handlerKembali() {
      this.konfirmasiHapus = false;
      this.iddivisi = "";
      this.namaDivisi = "";
    },

    DeleteDivisi() {
      axios.delete(`/divisi/${this.iddivisi}`)
        .then(() => {
          this.divisis = this.divisis.filter(divisi => divisi.id_divisi !== this.iddivisi);
          this.handlerKembali();
        })
        .catch(() => {
          this.message = "Gagal menghapus Divisi";
          setTimeout(() => {
            this.message = "";
          }, 2000)
        });
    },

    BukaEditDivisi(divisi) {
      this.formEdit = true;
      this.editDivisi = divisi
    },

    submitDivisi() {
      const data = {
        divisi: this.editDivisi.divisi,
      };

      axios.patch(`/divisi/${this.editDivisi.id_divisi}`, data)
      .then(() => {
        window.location.reload();
      })
      .catch(() => {
          this.message = "Gagal mengubah Divisi";
          setTimeout(() => {
            this.message = "";
          }, 2000)
        });
    }
  }
}
</script>

<style scoped>
td span {
  cursor: pointer;
}
</style>

<style scoped>
.absolute {
  position: fixed;
  background-color: var(--white);
  width: 100%;
  height: 100vh;
  top: 0;
}
</style>