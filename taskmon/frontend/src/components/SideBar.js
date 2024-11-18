import React, { useState } from "react";
import { HiOutlineQueueList } from "react-icons/hi2";
import { Link } from "react-router-dom";

function SideBar() {
  const [selected, setSelected] = useState(-1);
  return (
    <div className="w-56 h-[100vh] mt-5">
      <h1 className="text-blue-600 font-semibold ml-5 font-mono text-3xl">
        TaskFlow
      </h1>
      {selected === 0 ? (
        <div className="flex items-center bg-blue-100 p-1 rounded-r-full items-cetner mt-5">
          <HiOutlineQueueList size={28} className="text-blue-600 ml-5" />
          <h1 className="text-2xl font-mono ml-2 text-blue-600">Queues</h1>
        </div>
      ) : (
        <Link to={"/Queues"}>
          <div
            onClick={() => setSelected(0)}
            className="flex p-1 items-center ml-5 mt-5"
          >
            <HiOutlineQueueList size={28} />
            <h1 className="text-2xl font-mono ml-2">Queues</h1>
          </div>
        </Link>
      )}
    </div>
  );
}

export default SideBar;
