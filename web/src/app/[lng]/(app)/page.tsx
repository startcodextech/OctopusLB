export default function Home({params: {lng}}: {params: {lng: string}}) {
  return (
    <>
      Home {lng}
        <button>click</button>
    </>
  );
}
