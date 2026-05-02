import React from "react";

import styles from "./Chat.module.scss";

import clip from "@images/clip.png";
import arrow from "@images/arrow.png";
import microphone from "@images/microphone.png";

import call from "@images/call.png";
import settings from "@images/settigns.png";
// import { useRequest } from "@hooks/useRequest";

interface IMessage {}

interface RequestBody {
  Chat_id: number;
}

const requestBody: RequestBody = {
  Chat_id: 1,
};

const Chat = () => {
  // const [messages, isLoading, error] = useRequest<IMessage>(
  //   "http://26.132.220.182:8080/api/getallmessages?chat_id=1",
  // );
  let messages;

  console.log(messages);

  return (
    <div className={styles.chat}>
      <div className={styles.chatBody}>
        <div className={styles.header}>
          <div className={styles.headerLeft}>
            <div className={styles.avatarBody}>
              <img
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTTj5JRDSwFphGZvkw2rBcDkdzYi1bsGUyPfw&s"
                alt="avata"
              />
              <div className={styles.isOnline}></div>
            </div>

            <div className={styles.text}>
              <h3>Тимошка228</h3>
              <p>На связи</p>
            </div>
          </div>

          <div className={styles.headerRight}>
            <button className={styles.headerButton}>
              <img src={call} alt="call" />
            </button>

            <button className={styles.headerButton}>
              <img src={settings} alt="settings" />
            </button>
          </div>
        </div>

        <div className={styles.messages}>
          {new Array(100).fill(0).map((elem, ind) => (
            <div key={Date.now() + ind}>{ind + 1}</div>
          ))}
        </div>

        <div className={styles.sender}>
          <div className={styles.left}>
            <button className={styles.btn}>
              <img src={clip} alt="clip" className={styles.invert} />
            </button>

            <input
              type="text"
              className={styles.input}
              placeholder="Сообщение..."
            />
          </div>

          <div className={styles.right}>
            <button className={styles.btn}>
              <img
                src={microphone}
                alt="microphone"
                className={styles.invert}
              />
            </button>

            <button className={styles.btn}>
              <img src={arrow} alt="arrow" className={styles.invert} />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Chat;
