import { useState } from "react";

export const Favicon = ({ url }: { url: string }) => {
  const src = `${new URL(url).origin}/favicon.ico`;
  const [isLoading, setIsLoading] = useState(true);

  return (
    <img
      src={src}
      alt="favicon"
      onLoad={() => setIsLoading(false)}
      style={
        isLoading
          ? { display: "none" }
          : {
              width: 24,
              height: 24,
            }
      }
    />
  );
};
