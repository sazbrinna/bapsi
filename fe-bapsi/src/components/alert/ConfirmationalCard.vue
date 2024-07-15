<template>
  <section class="confirm full">
    <div class="card">
      <div class="content line">
        <p>Yakin ingin menghapus <strong>{{ nama }}</strong>?</p>
      </div>
      <div class="content">
        <span>
          <a href="#" class="green" @click="kembali">Kembali</a>
          <a href="#" class="red" @click="hapus">Hapus</a>
        </span>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  props: {
    nama: String,
  },
  
  methods: {
    kembali() {
      this.$emit("kembali");
    },
    hapus() {
      this.$emit("hapus");
    }
  }
}
</script>

<style scoped>
section.confirm {
  position: fixed;
  background-color: #ffffffcb;
  top: 0;
}

.card {
  box-sizing: border-box;
  width: 420px;
  max-width: 95%;
  border: 2px solid var(--red);
  border-radius: 8px;
  background-color: var(--white);
  box-shadow: 0 2px 7px 0 var(--darkgray);
}

.content {
  box-sizing: border-box;
  padding: 1rem;
}

.content span {
  display: flex;
  align-items: center;
  justify-content: space-around;
}

.line {
  background-color: var(--red) !important;
  color: var(--white);
  border-bottom: 1px solid var(--darkgray);
}
</style>