<template>
  <div class="home">
    <div class="header">
      <div class="logo">
        <img class="desktop" src="../assets/logo.jpg" alt="" />
        <img class="mobile" src="../assets/moblogo.png" alt="" />
        <img class="bg" src="../assets/logo2.png" alt="" />
      </div>
      <h1 class="mobile">
        ЗАГРУЖЕННОСТЬ МФЦ
      </h1>
    </div>
    <div class="control">
      <h1 class="desktop">
        ЗАГРУЖЕННОСТЬ <br />
        МФЦ
      </h1>
      <h2>
        здесь вы можете узнать, <br />
        какое <span class="orange">отделение мфц будет</span> <br />
        наиболее удобным для вас <br />
        в данный момент
      </h2>
      <div class="left input-label">
        <div class="circle"></div>
        <div class="num">
          1
        </div>
        <h3 class="n1">откуда вы планируете добираться?</h3>
      </div>
      <div class="input-block">
        <div class="search-icon">
          <svg
            width="46"
            height="42"
            viewBox="0 0 46 42"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle
              cx="18.456"
              cy="18.0759"
              r="15.5759"
              stroke="#E75138"
              stroke-width="5"
            />
            <line
              x1="31.92"
              y1="28.0043"
              x2="44.1478"
              y2="40.2322"
              stroke="#E75138"
              stroke-width="5"
            />
          </svg>
        </div>
        <input
          type="text"
          name="Address"
          id=""
          placeholder="Ваше местоположение"
        />
      </div>
      <div class="rigth input-label">
        <div class="circle"></div>
        <div class="num">
          2
        </div>
        <h3 class="n2">выбор способа передвижения</h3>
      </div>
      <div class="transport-pick">
        <div class="item">
          <svg
            width="106"
            height="106"
            viewBox="0 0 106 106"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M53 23.1875C59.4031 23.1875 64.5938 17.9968 64.5938 11.5938C64.5938 5.1907 59.4031 0 53 0C46.5969 0 41.4062 5.1907 41.4062 11.5938C41.4062 17.9968 46.5969 23.1875 53 23.1875Z"
              fill="white"
            />
            <path
              d="M62.9375 26.5H43.0625C39.5516 26.5104 36.1874 27.9097 33.7048 30.3923C31.2222 32.8749 29.8229 36.2391 29.8125 39.75V62.01C29.8125 64.2563 31.5578 66.1672 33.802 66.2479C34.3582 66.2682 34.9127 66.1762 35.4325 65.9774C35.9523 65.7785 36.4267 65.477 36.8273 65.0906C37.2279 64.7043 37.5465 64.2412 37.7641 63.729C37.9817 63.2168 38.0938 62.6659 38.0937 62.1094V41.4621C38.0885 41.0349 38.2449 40.6214 38.5318 40.3046C38.8186 39.9878 39.2146 39.7912 39.6403 39.7541C39.8668 39.7391 40.0941 39.7708 40.3078 39.8473C40.5216 39.9237 40.7174 40.0434 40.883 40.1987C41.0486 40.354 41.1805 40.5417 41.2705 40.7502C41.3606 40.9586 41.4067 41.1833 41.4062 41.4104V101.238C41.4062 102.501 41.9079 103.712 42.8009 104.605C43.6939 105.498 44.9051 106 46.168 106C47.4309 106 48.642 105.498 49.535 104.605C50.428 103.712 50.9297 102.501 50.9297 101.238V71.7032C50.9224 71.1686 51.1173 70.6509 51.4755 70.2539C51.8337 69.8569 52.3287 69.61 52.8613 69.5625C53.1446 69.5435 53.4288 69.5829 53.6962 69.6784C53.9636 69.7739 54.2086 69.9234 54.4157 70.1176C54.6229 70.3117 54.788 70.5465 54.9006 70.8071C55.0132 71.0678 55.0709 71.3489 55.0703 71.6328V101.238C55.0703 102.501 55.572 103.712 56.465 104.605C57.358 105.498 58.5691 106 59.832 106C61.0949 106 62.3061 105.498 63.1991 104.605C64.0921 103.712 64.5937 102.501 64.5937 101.238V41.4621C64.5885 41.0349 64.7449 40.6214 65.0318 40.3046C65.3186 39.9878 65.7146 39.7912 66.1403 39.7541C66.3668 39.7391 66.594 39.7708 66.8078 39.8473C67.0216 39.9237 67.2174 40.0434 67.383 40.1987C67.5486 40.354 67.6805 40.5417 67.7705 40.7502C67.8606 40.9586 67.9067 41.1833 67.9062 41.4104V62.0141C67.9062 64.2604 69.6515 66.1713 71.8957 66.2521C72.4522 66.2724 73.0071 66.1803 73.5272 65.9812C74.0473 65.7821 74.5219 65.4801 74.9225 65.0933C75.3232 64.7066 75.6418 64.243 75.8591 63.7303C76.0765 63.2176 76.1881 62.6663 76.1875 62.1094V39.75C76.1771 36.2391 74.7778 32.8749 72.2952 30.3923C69.8126 27.9097 66.4484 26.5104 62.9375 26.5Z"
              fill="white"
            />
          </svg>
        </div>
        <div class="item">
          <svg
            width="121"
            height="121"
            viewBox="0 0 121 121"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M104.726 51.2032L97.8285 30.5071C96.828 27.4946 94.9036 24.8741 92.3286 23.0178C89.7537 21.1615 86.6592 20.1638 83.485 20.1667H37.515C34.3408 20.1638 31.2463 21.1615 28.6714 23.0178C26.0965 24.8741 24.172 27.4946 23.1715 30.5071L16.2745 51.2032C14.4424 51.9719 12.8778 53.264 11.7766 54.9177C10.6753 56.5714 10.0863 58.5132 10.0833 60.5V85.7083C10.0833 89.5047 12.216 92.7767 15.3216 94.496C15.2561 94.8287 15.125 95.1312 15.125 95.4791V105.875C15.125 107.212 15.6562 108.495 16.6017 109.44C17.5472 110.385 18.8295 110.917 20.1667 110.917H25.2083C26.5455 110.917 27.8278 110.385 28.7733 109.44C29.7188 108.495 30.25 107.212 30.25 105.875V95.7917H90.75V105.875C90.75 107.212 91.2812 108.495 92.2267 109.44C93.1722 110.385 94.4545 110.917 95.7917 110.917H100.833C102.17 110.917 103.453 110.385 104.398 109.44C105.344 108.495 105.875 107.212 105.875 105.875V95.4791C105.875 95.1312 105.744 94.8237 105.678 94.496C107.258 93.6345 108.578 92.3645 109.5 90.8187C110.421 89.2729 110.91 87.508 110.917 85.7083V60.5C110.917 56.3205 108.355 52.7308 104.726 51.2032ZM20.1667 85.7083V60.5H100.833L100.843 85.7083H20.1667ZM37.515 30.25H83.4799C85.6529 30.25 87.5788 31.6314 88.2645 33.6985L93.8406 50.4167H60.5H27.1595L32.7305 33.6985C33.065 32.6942 33.7072 31.8206 34.5659 31.2016C35.4247 30.5827 36.4565 30.2497 37.515 30.25Z"
              fill="white"
            />
            <path
              d="M32.7708 80.6667C36.9475 80.6667 40.3333 77.2808 40.3333 73.1042C40.3333 68.9275 36.9475 65.5417 32.7708 65.5417C28.5942 65.5417 25.2083 68.9275 25.2083 73.1042C25.2083 77.2808 28.5942 80.6667 32.7708 80.6667Z"
              fill="white"
            />
            <path
              d="M88.2292 80.6667C92.4058 80.6667 95.7917 77.2808 95.7917 73.1042C95.7917 68.9275 92.4058 65.5417 88.2292 65.5417C84.0525 65.5417 80.6667 68.9275 80.6667 73.1042C80.6667 77.2808 84.0525 80.6667 88.2292 80.6667Z"
              fill="white"
            />
          </svg>
        </div>
        <div class="item">
          <svg
            width="85"
            height="103"
            viewBox="0 0 85 103"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M25.7917 78.84C29.8648 78.84 33.1667 75.4882 33.1667 71.3535C33.1667 67.2189 29.8648 63.8671 25.7917 63.8671C21.7186 63.8671 18.4167 67.2189 18.4167 71.3535C18.4167 75.4882 21.7186 78.84 25.7917 78.84Z"
              fill="white"
            />
            <path
              d="M60.2083 78.84C64.2814 78.84 67.5833 75.4882 67.5833 71.3535C67.5833 67.2189 64.2814 63.8671 60.2083 63.8671C56.1352 63.8671 52.8333 67.2189 52.8333 71.3535C52.8333 75.4882 56.1352 78.84 60.2083 78.84Z"
              fill="white"
            />
            <path
              d="M42.5 0C19.0187 0 0 2.69805 0 21.5844V75.5453C0 80.2939 2.07187 84.5568 5.3125 87.5246V97.1297C5.3125 100.098 7.70312 102.526 10.625 102.526H15.9375C18.8594 102.526 21.25 100.098 21.25 97.1297V91.7336H63.75V97.1297C63.75 100.098 66.1406 102.526 69.0625 102.526H74.375C77.2969 102.526 79.6875 100.098 79.6875 97.1297V87.5246C82.9281 84.5568 85 80.2939 85 75.5453V21.5844C85 2.69805 65.9813 0 42.5 0ZM72.5687 16.1343H12.4313C15.3531 13.2744 22.8969 10.7922 42.5 10.7922C62.1031 10.7922 69.6469 13.2744 72.5687 16.1343ZM74.375 26.9265V43.1687H10.625V26.9265H74.375ZM72.5687 79.4844L71.0281 80.9414H13.9719L12.4313 79.4844C11.8695 78.998 11.4174 78.3944 11.1057 77.7147C10.794 77.035 10.6301 76.2951 10.625 75.5453V53.9609H74.375V75.5453C74.375 77.5418 73.2594 78.8909 72.5687 79.4844Z"
              fill="white"
            />
          </svg>
        </div>
      </div>
      <div class="left input-label">
        <div class="circle"></div>
        <div class="num">
          3
        </div>
        <h3 class="n3">выбор услуги</h3>
      </div>
      <div class="input-block">
        <div class="search-icon">
          <svg
            width="46"
            height="42"
            viewBox="0 0 46 42"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle
              cx="18.456"
              cy="18.0759"
              r="15.5759"
              stroke="#E75138"
              stroke-width="5"
            />
            <line
              x1="31.92"
              y1="28.0043"
              x2="44.1478"
              y2="40.2322"
              stroke="#E75138"
              stroke-width="5"
            />
          </svg>
        </div>
        <input
          type="text"
          name="Address"
          id=""
          placeholder="Ваше местоположение"
        />
      </div>
    </div>

    <div class="map-cnt">
      <div id="map"></div>
    </div>

    <div class="best-offer">
      <div class="header">
        наиболее удобное отделение
      </div>
      <div class="icon">
        <svg
          width="261"
          height="377"
          viewBox="0 0 261 377"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M131.083 343.89C130.537 352.867 131.288 347.6 130.674 376.607C65.1288 318.444 10.9057 255.014 1.08415 163.096C-5.25892 103.375 16.2939 53.8176 64.1057 20.5074C83.8852 6.70858 105.779 -1.00691 129.787 0.1059C144.247 21.8428 138.313 45.5828 135.721 69.026C131.902 80.6734 121.671 83.4925 113.077 88.6115C84.9764 105.378 73.9272 131.047 79.9293 165.544C85.4539 197.593 100.8 224.226 118.738 249.376C125.286 258.575 133.811 266.364 136.471 278.457C138.995 300.787 142.269 323.044 131.083 343.89Z"
            fill="#E75138"
          />
          <path
            d="M131.083 343.89C130.674 319.556 130.265 295.297 129.855 270.963C119.761 254.271 123.922 240.398 136.471 227.712C149.84 214.136 163.617 201.005 172.074 182.754C191.376 141.358 182.714 109.828 145.543 87.4236C139.813 83.9368 133.129 83.7142 129.787 76.5922C129.787 51.6653 129.719 25.032 129.719 0.10498C187.216 1.51454 228.139 31.8572 249.896 88.5364C274.723 153.154 255.557 211.391 215.794 260.726C190.967 291.662 159.593 316.366 131.083 343.89Z"
            fill="#CE9966"
          />
          <path
            d="M112.054 80.1542C123.376 75.5546 135.926 75.332 147.384 79.4865C168.255 86.9052 184.351 101.372 188.648 127.56C197.447 181.346 164.435 213.469 129.924 244.034C129.924 253.085 129.924 262.061 129.924 271.038C95.0026 239.508 72.9724 199.373 70.7216 150.113C69.0847 114.429 84.5672 91.2823 112.054 80.1542Z"
            fill="#FEFEFE"
          />
          <path
            d="M129.719 161.76C140.417 161.76 149.089 152.327 149.089 140.691C149.089 129.055 140.417 119.622 129.719 119.622C119.021 119.622 110.349 129.055 110.349 140.691C110.349 152.327 119.021 161.76 129.719 161.76Z"
            fill="#CE9966"
          />
        </svg>
      </div>
      <div class="block">
        <p class="address">Партизанский просп., 28А</p>
        <p class="number">+7 (423) 201-01-56</p>
        <p class="road">дорога до отделения и получение услуги займут около</p>
        <p class="time">~15 минут</p>
      </div>
      <div class="stat-header">Статистика загруженности:</div>
      <div class="stat">
        <div class="item">
          <p>в очереди:</p>
          <div class="icon">
            <svg
              width="106"
              height="106"
              viewBox="0 0 106 106"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M53 23.1875C59.4031 23.1875 64.5938 17.9968 64.5938 11.5938C64.5938 5.1907 59.4031 0 53 0C46.5969 0 41.4062 5.1907 41.4062 11.5938C41.4062 17.9968 46.5969 23.1875 53 23.1875Z"
                fill="#783424"
              />
              <path
                d="M62.9375 26.5H43.0625C39.5516 26.5104 36.1874 27.9097 33.7048 30.3923C31.2222 32.8749 29.8229 36.2391 29.8125 39.75V62.01C29.8125 64.2563 31.5578 66.1672 33.802 66.2479C34.3582 66.2682 34.9127 66.1762 35.4325 65.9774C35.9523 65.7785 36.4267 65.477 36.8273 65.0906C37.2279 64.7043 37.5465 64.2412 37.7641 63.729C37.9817 63.2168 38.0938 62.6659 38.0937 62.1094V41.4621C38.0885 41.0349 38.2449 40.6214 38.5318 40.3046C38.8186 39.9878 39.2146 39.7912 39.6403 39.7541C39.8668 39.7391 40.0941 39.7708 40.3078 39.8473C40.5216 39.9237 40.7174 40.0434 40.883 40.1987C41.0486 40.354 41.1805 40.5417 41.2705 40.7502C41.3606 40.9586 41.4067 41.1833 41.4062 41.4104V101.238C41.4062 102.501 41.9079 103.712 42.8009 104.605C43.6939 105.498 44.9051 106 46.168 106C47.4309 106 48.642 105.498 49.535 104.605C50.428 103.712 50.9297 102.501 50.9297 101.238V71.7032C50.9224 71.1686 51.1173 70.6509 51.4755 70.2539C51.8337 69.8569 52.3287 69.61 52.8613 69.5625C53.1446 69.5435 53.4288 69.5829 53.6962 69.6784C53.9636 69.7739 54.2086 69.9234 54.4157 70.1176C54.6229 70.3117 54.788 70.5465 54.9006 70.8071C55.0132 71.0678 55.0709 71.3489 55.0703 71.6328V101.238C55.0703 102.501 55.572 103.712 56.465 104.605C57.358 105.498 58.5691 106 59.832 106C61.0949 106 62.3061 105.498 63.1991 104.605C64.0921 103.712 64.5937 102.501 64.5937 101.238V41.4621C64.5885 41.0349 64.7449 40.6214 65.0318 40.3046C65.3186 39.9878 65.7146 39.7912 66.1403 39.7541C66.3668 39.7391 66.594 39.7708 66.8078 39.8473C67.0216 39.9237 67.2174 40.0434 67.383 40.1987C67.5486 40.354 67.6805 40.5417 67.7705 40.7502C67.8606 40.9586 67.9067 41.1833 67.9062 41.4104V62.0141C67.9062 64.2604 69.6515 66.1713 71.8957 66.2521C72.4522 66.2724 73.0071 66.1803 73.5272 65.9812C74.0473 65.7821 74.5219 65.4801 74.9225 65.0933C75.3232 64.7066 75.6418 64.243 75.8591 63.7303C76.0765 63.2176 76.1881 62.6663 76.1875 62.1094V39.75C76.1771 36.2391 74.7778 32.8749 72.2952 30.3923C69.8126 27.9097 66.4484 26.5104 62.9375 26.5Z"
                fill="#783424"
              />
            </svg>
          </div>
          <span>3</span>
        </div>
        <div class="item">
          <p>обслужено сегодня:</p>
          <div class="icon">
            <svg
              width="106"
              height="106"
              viewBox="0 0 106 106"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M53 23.1875C59.4031 23.1875 64.5938 17.9968 64.5938 11.5938C64.5938 5.1907 59.4031 0 53 0C46.5969 0 41.4062 5.1907 41.4062 11.5938C41.4062 17.9968 46.5969 23.1875 53 23.1875Z"
                fill="#783424"
              />
              <path
                d="M62.9375 26.5H43.0625C39.5516 26.5104 36.1874 27.9097 33.7048 30.3923C31.2222 32.8749 29.8229 36.2391 29.8125 39.75V62.01C29.8125 64.2563 31.5578 66.1672 33.802 66.2479C34.3582 66.2682 34.9127 66.1762 35.4325 65.9774C35.9523 65.7785 36.4267 65.477 36.8273 65.0906C37.2279 64.7043 37.5465 64.2412 37.7641 63.729C37.9817 63.2168 38.0938 62.6659 38.0937 62.1094V41.4621C38.0885 41.0349 38.2449 40.6214 38.5318 40.3046C38.8186 39.9878 39.2146 39.7912 39.6403 39.7541C39.8668 39.7391 40.0941 39.7708 40.3078 39.8473C40.5216 39.9237 40.7174 40.0434 40.883 40.1987C41.0486 40.354 41.1805 40.5417 41.2705 40.7502C41.3606 40.9586 41.4067 41.1833 41.4062 41.4104V101.238C41.4062 102.501 41.9079 103.712 42.8009 104.605C43.6939 105.498 44.9051 106 46.168 106C47.4309 106 48.642 105.498 49.535 104.605C50.428 103.712 50.9297 102.501 50.9297 101.238V71.7032C50.9224 71.1686 51.1173 70.6509 51.4755 70.2539C51.8337 69.8569 52.3287 69.61 52.8613 69.5625C53.1446 69.5435 53.4288 69.5829 53.6962 69.6784C53.9636 69.7739 54.2086 69.9234 54.4157 70.1176C54.6229 70.3117 54.788 70.5465 54.9006 70.8071C55.0132 71.0678 55.0709 71.3489 55.0703 71.6328V101.238C55.0703 102.501 55.572 103.712 56.465 104.605C57.358 105.498 58.5691 106 59.832 106C61.0949 106 62.3061 105.498 63.1991 104.605C64.0921 103.712 64.5937 102.501 64.5937 101.238V41.4621C64.5885 41.0349 64.7449 40.6214 65.0318 40.3046C65.3186 39.9878 65.7146 39.7912 66.1403 39.7541C66.3668 39.7391 66.594 39.7708 66.8078 39.8473C67.0216 39.9237 67.2174 40.0434 67.383 40.1987C67.5486 40.354 67.6805 40.5417 67.7705 40.7502C67.8606 40.9586 67.9067 41.1833 67.9062 41.4104V62.0141C67.9062 64.2604 69.6515 66.1713 71.8957 66.2521C72.4522 66.2724 73.0071 66.1803 73.5272 65.9812C74.0473 65.7821 74.5219 65.4801 74.9225 65.0933C75.3232 64.7066 75.6418 64.243 75.8591 63.7303C76.0765 63.2176 76.1881 62.6663 76.1875 62.1094V39.75C76.1771 36.2391 74.7778 32.8749 72.2952 30.3923C69.8126 27.9097 66.4484 26.5104 62.9375 26.5Z"
                fill="#783424"
              />
            </svg>
          </div>
          <span>253</span>
        </div>
      </div>
    </div>
    <div class="stats">
      <info-card :address="'ул.Хуя'" :time="'дохуя'"></info-card>
      <info-card :address="'ул.Хуя'" :time="'дохуя'"></info-card>
      <info-card :address="'ул.Хуя'" :time="'дохуя'"></info-card>
      <info-card :address="'ул.Хуя'" :time="'дохуя'"></info-card>
    </div>
    <div class="bad-version">
      <h3>
        если вы не нашли для себя подходящего варианта, можно перенести
        посещение мфц
      </h3>
      <input
        class="btn"
        type="button"
        value="записаться на более удобное время"
      />
    </div>
  </div>
</template>

<script>
import GoogleMapsLoader from "google-maps";
import InfoCard from "../components/InfoCard";

export default {
  name: "home",
  components: {
    InfoCard
  },
  data() {
    return {
      map: undefined,
      googleMapsClient: require("@google/maps").createClient({
        key: "AIzaSyC0zt4dGxQo4j_dt9z8dofi1UHQOApc8S0"
      })
    };
  },
  mounted() {
    GoogleMapsLoader.KEY = "AIzaSyC0zt4dGxQo4j_dt9z8dofi1UHQOApc8S0";
    GoogleMapsLoader.LIBRARIES = ["directions", "places"];
    GoogleMapsLoader.VERSION = "3";
    GoogleMapsLoader.REGION = "ru";

    GoogleMapsLoader.load(google => {
      let chicago = new google.maps.LatLng(41.850033, -87.6500523);

      this.map = new google.maps.Map(document.getElementById("map"), {
        zoom: 7,
        center: chicago
      });
    });
  }
};
</script>

<style lang="scss">
.n1 {
  @media (max-width: 600px) {
    padding-left: 20%;
  }
}

.n2 {
  @media (max-width: 600px) {
    padding-left: 20%;
  }
}

.n3 {
  @media (max-width: 600px) {
    padding-left: 20%;
  }
}

.left {
}

.rigth {
}

.mobile {
  display: none;

  @media (max-width: 600px) {
    display: inline;
  }
}

.desktop {
  display: inline;

  @media (max-width: 600px) {
    display: none;
  }
}

.home {
  font-family: Circe;
  display: grid;
  grid-auto-rows: 200px;
  grid-row-gap: 0px;
  grid-template-columns: repeat(12, 1fr);
  grid-template-areas:
    ". . lg lg lg lg lg lg lg lg . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . ct ct ct ct ct ct ct ct . ."
    ". . mp mp mp mp mp mp mp mp . ."
    ". . mp mp mp mp mp mp mp mp . ."
    ". . mp mp mp mp mp mp mp mp . ."
    ". . mp mp mp mp mp mp mp mp . ."
    ". . mp mp mp mp mp mp mp mp . ."
    ". . in in in in in in in in . ."
    ". . in in in in in in in in . ."
    ". . in in in in in in in in . ."
    ". . in in in in in in in in . ."
    ". . st st st st st st st st . ."
    ". . st st st st st st st st . ."
    ". . bd bd bd bd bd bd bd bd . ."
    ". . bd bd bd bd bd bd bd bd . .";
  position: relative;

  @media (max-width: 600px) {
    grid-auto-rows: 100px;
    grid-template-areas:
      "lg lg lg lg lg lg lg lg lg lg lg lg"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "ct ct ct ct ct ct ct ct ct ct ct ct"
      "mp mp mp mp mp mp mp mp mp mp mp mp"
      "mp mp mp mp mp mp mp mp mp mp mp mp"
      "mp mp mp mp mp mp mp mp mp mp mp mp"
      "mp mp mp mp mp mp mp mp mp mp mp mp"
      ". . in in in in in in in in . ."
      ". . in in in in in in in in . ."
      ". . in in in in in in in in . ."
      ". . in in in in in in in in . ."
      ". . st st st st st st st st . ."
      ". . st st st st st st st st . ."
      ". . bd bd bd bd bd bd bd bd . ."
      ". . bd bd bd bd bd bd bd bd . .";
  }
}

h1 {
  font-style: normal;
  font-weight: 800;
  font-size: 90px;
  line-height: 75px;
  text-align: left;
  color: #55220d;

  @media (max-width: 600px) {
    grid-area: lg;
    font-style: normal;
    font-weight: 800;
    font-size: 22px;
    line-height: 34px;
  }
}

h2 {
  font-style: normal;
  font-weight: 800;
  font-size: 50px;
  text-align: left;
  line-height: 60px;
  margin-bottom: 200px;

  @media (max-width: 600px) {
    margin-bottom: 100px;
    font-style: normal;
    justify-self: center;
    padding-left: 15%;
    font-weight: 800;
    font-size: 13px;
    line-height: 16px;
  }
}

h3 {
  font-style: normal;
  font-weight: 800;
  font-size: 36px;
  line-height: 53px;

  @media (max-width: 600px) {
    font-style: normal;
    font-weight: 800;
    font-size: 20px;
    line-height: 20px;
    text-align: left;
    width: 40%;
  }
}

.num {
  font-weight: 800;
  font-size: 144px;
  line-height: 212px;
}

.control {
  grid-area: ct;
  text-transform: uppercase;
}

.map-cnt {
  width: 100%;
  position: relative;
  grid-area: mp;
}

.header {
  grid-area: lg;
  display: flex;
  justify-content: start;

  .bg {
    position: absolute;
    top: -50px;
    right: -180px;

    @media (max-width: 600px) {
      display: none;
    }
  }

  .logo {
    height: 150px;
    width: 200px;
    @media (max-width: 600px) {
      padding-top: 16px;
      width: 100px;
      height: 100px;
    }
  }
}

.input-label {
  position: relative;
  margin-top: 32px;

  .circle {
    position: absolute;
    width: 52px;
    height: 52px;
    background: #e75138;
    left: 140px;
    top: 5px;
    border-radius: 50%;
    z-index: 999;
  }

  .num {
    z-index: 1000;
    position: absolute;
    left: 130px;
    top: -75px;
  }
}

.transport-pick {
  display: flex;
  justify-content: center;
  align-items: center;

  .item {
    width: 145px;
    height: 145px;
    border-radius: 50%;
    background: #e75138;
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 0 32px;
  }
}

.input-block {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  height: 90px;
  width: 60%;
  border: 4px solid #e75138;
  box-sizing: border-box;
  border-radius: 80px;
  position: relative;
  padding: 0 16px;
  margin: 0 auto;

  .search-icon {
    width: 44px;
    height: 44px;
  }

  input {
    width: calc(100% - 60px);
    margin-left: 16px;
    height: 60px;
    border: none;
  }
}

.best-offer {
  display: grid;
  position: relative;
  grid-area: in;
  grid-template-rows: 80px 200px 200px 50px 250px;
  grid-template-areas:
    ".  .  hd hd hd hd"
    "ic ic of of of of"
    "ic ic of of of of"
    ".  .  hs hs hs hs"
    ".  .  s  s  s  s";

  .icon {
    grid-area: ic;
  }

  .header {
    grid-area: hd;
    font-weight: 800;
    font-size: 40px;
    line-height: 59px;
  }

  .block {
    padding: 0 24px;
    text-align: left;
    grid-area: of;
    background: #eb4b34;
    border-radius: 20px;

    .address {
      font-weight: 800;
      font-size: 42px;
      line-height: 62px;
    }

    .number {
      font-weight: bold;
      font-size: 36px;
      line-height: 53px;
    }

    .road {
      font-style: normal;
      font-weight: normal;
      font-size: 24px;
      line-height: 28px;
    }

    .time {
      font-style: normal;
      font-weight: 800;
      font-size: 72px;
      line-height: 14px;
    }
  }

  .stat-header {
    font-weight: 800;
    font-size: 36px;
    line-height: 53px;
    grid-area: hs;
    text-align: left;
    display: flex;
    align-items: center;
  }

  .stat {
    grid-area: s;
    display: flex;

    .item {
      margin-right: 64px;
      position: relative;
      p {
        font-style: normal;
        font-weight: normal;
        font-size: 24px;
        line-height: 35px;
      }

      .icon {
        display: inline;
      }

      span {
        font-weight: 800;
        font-size: 72px;
        line-height: 106px;
        text-align: center;
        margin-bottom: 25%;
      }
    }
  }
}

.stats {
  grid-area: st;
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  position: relative;
  flex-wrap: wrap;
}

.bad-version {
  grid-area: bd;

  .btn {
    height: 90px;
    width: 60%;
    background: #eb4b34;
    border: 4px solid #eb4b34;
    box-sizing: border-box;
    border-radius: 80px;
    font-style: normal;
    font-weight: 800;
    font-size: 24px;
    line-height: 35px;
  }
}

#map {
  height: 100%;
}
</style>
