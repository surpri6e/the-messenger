import { useUserStore } from "@stores/useUserStore";
import styles from "./AvatarSection.module.scss";
import basicAvatar from "@images/user.png";
import HeaderText from "@components/HeaderText/HeaderText";
import camera from "@images/camera.png";
import { useAvatarChange } from "@hooks/useAvatarChange";
import Loader from "@components/Loader/Loader";

const AvatarSection = () => {
  const { user } = useUserStore((state) => state);

  const { avatarError, fn: changeAvatar, isLoading } = useAvatarChange();

  const onChange = async (
    e: React.ChangeEvent<HTMLInputElement, HTMLInputElement>,
  ) => {
    console.log(e.target.files);
    if (e.target.files && e.target.files[0])
      await changeAvatar(e.target.files[0]);
  };

  return (
    <section className={styles.avatarSection}>
      <div className={styles.avatarSectionBody}>
        <label
          className={`${styles.avatarBody} ${avatarError ? styles.avatarBodyError : ""}`}
          htmlFor="avatarChange"
        >
          {isLoading ? (
            <Loader />
          ) : user?.avatar_link.length === 0 ? (
            <img
              src={basicAvatar}
              alt="Ваш базовый аватар"
              className={styles.basicAvatar}
            />
          ) : (
            <img
              src={user?.avatar_link}
              alt="Ваш аватар"
              className={styles.avatar}
            />
          )}

          <span className={styles.cameraBody}>
            <img
              src={camera}
              alt="Прикрепить ваше фото"
              className={styles.camera}
            />
          </span>

          <input
            type="file"
            id="avatarChange"
            accept=".png, .jpg, .jpeg"
            onChange={onChange}
          />
        </label>

        <div>
          <HeaderText
            text={
              user!.username.length > 12
                ? `${user!.username.slice(0, 9)}...`
                : user!.username
            }
            isDarken
          />
          <p className={styles.isOnline}>На связи</p>
        </div>
      </div>
    </section>
  );
};

export default AvatarSection;
