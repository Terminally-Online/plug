import Viewport from "./components/Viewport";

const getCanvas = async (id: string) => {
  return { id: id };

  // const res = await fetch(`http://localhost:3000/api/canvas/${id}`);
  //
  // const json = await res.json();
  //
  // return json;
}

export default async function Page({ params }: { params: { id: string } }) { 
  const canvas = await getCanvas(params.id);

  return <Viewport />
}
