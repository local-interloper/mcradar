import { Outlet } from "react-router";

const CoreLayout = () => {
  return (
    <article className="fixed top-0 left-0 w-full h-full">
      <Outlet />
    </article>
  );
};

export default CoreLayout;
