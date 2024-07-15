<template>
  <div class="alert">
    <div class="wrapperCard">
      <div class="status">
        <p>Upps ada yang salah!</p>
      </div>
      <div class="message">
        <p>{{ msg }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      shown: true,
    }
  },

  props: {
    msg: String,
  },
}
</script>

<style scoped>
.alert {
  position: fixed;
  right: 0;
  bottom: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: 1s;
}

.wrapperCard {
  display: flex;
  flex-direction: column;
  width: 360px;
  max-width: 90%;
  border-style: solid;
  border-radius: 8px;
  overflow: hidden;
}

.alert .status {
  background-color: var(--yellow);
  width: 100%;
  padding: 0.5rem 1rem;
  font-weight: 600;
  box-sizing: border-box;
  color: var(--black);
}

.alert .message {
  box-sizing: border-box;
  width: 100%;
  padding: 0.5rem 1rem;
  background-color: white;
}

@media screen and (max-width: 768px) {
  .alert {
    right: auto;
    left: auto;
    width: 100%;
  }
}
</style>