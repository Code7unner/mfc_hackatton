<template>
  <div class="info-card">
    <div class="icon">
      <svg
        width="70"
        height="97"
        viewBox="0 0 70 97"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
      >
        <rect width="70" height="97" fill="url(#pattern0)" />
        <defs>
          <pattern
            id="pattern0"
            patternContentUnits="objectBoundingBox"
            width="1"
            height="1"
          >
            <use
              xlink:href="#image0"
              transform="translate(-0.00559846) scale(0.0187259 0.0135135)"
            />
          </pattern>
          <image
            id="image0"
            width="54"
            height="74"
            xlink:href="data:image/png;base64,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"
          />
        </defs>
      </svg>
    </div>
    <div class="block">
      <div class="address">{{ address }}</div>
      <div class="road">получение услуги займет примерно</div>
      <div class="time">~ {{ time }} минут</div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    address: { required: false },
    time: { required: false }
  }
};
</script>

<style lang="scss">
.info-card {
  width: 40%;
  display: grid;
  grid-template-areas:
    "ic ic bk bk bk bk"
    "ic ic bk bk bk bk"
    "ic ic bk bk bk bk"
    "ic ic bk bk bk bk";

  .icon {
    grid-area: ic;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .block {
    border: 4px solid #cb9b63;
    box-sizing: border-box;
    border-radius: 20px;
    text-align: left;
    padding-left: 16px;
    grid-auto-rows: 40px;
    margin-left: 16px;

    grid-area: bk;
    display: grid;
    grid-template-areas:
      "ad"
      "rd"
      "tm";

    .address {
      grid-area: ad;
      font-style: normal;
      font-weight: 800;
      font-size: 22px;
      line-height: 32px;
    }

    .road {
      font-style: normal;
      font-weight: normal;
      font-size: 14px;
      line-height: 14px;
      grid-area: rd;
    }

    .time {
      grid-area: tm;
      font-style: normal;
      font-weight: 800;
      font-size: 36px;
      line-height: 14px;
    }
  }
}
</style>
