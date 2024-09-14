import { Tabs } from '@modules/dashboard';
import {Navbar} from "@components/layouts/dashboard";

export default function Home({ params: { lng } }: { params: { lng: string } }) {
  return (
    <>
          <Tabs />
      Home {lng}
      <button>click</button>
    </>
  );
}
