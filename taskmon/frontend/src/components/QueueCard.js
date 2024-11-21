import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

const headings = ["Queue", "State", "Memory Usage"];

// const tableData = [];

function QueueCard() {
  const [tableData, setTableData] = useState([]);
  const fetchData = async () => {
    try {
      const resp = await axios.get("http://localhost:9005/Queues");
      // console.log(resp);
      // const result = resp.json();
      // console.log(result);
      setTableData(resp.data?.AllQueuesInformation);
      // console.log(resp.data?.AllQueuesInformation);
    } catch (err) {
      console.log(err);
    }
  };
  useEffect(() => {
    fetchData();
    const IntervalId = setInterval(fetchData, 5000);
    return () => clearInterval(IntervalId);
  }, []);
  return (
    <div>
      <table className="w-[50vw] rounded-md">
        <thead className="bg-gray-50 text-xs rounded-md p-2">
          <tr className="border-l border-r uppercase p-2 border-t border-gray-300">
            {headings.map((heading, idx) => {
              return (
                <th key={idx} scope="col" className=" px-6 py-3 font-mono">
                  {heading}
                </th>
              );
            })}
          </tr>
        </thead>
        <tbody className="border border-gray-300 rounded-md">
          {tableData?.map((row, idx) => {
            return (
              <tr
                key={idx}
                className="text-sm border-x hover:bg-gray-50 border-gray-300 border-t justify-center"
              >
                <Link to={`/Queues/${row.Name}`}>
                  {/* {row.map((data, idx) => {
                  return ( */}
                  <td key={idx} className="px-10 text-center py-4 font-mono">
                    {row?.Name}
                  </td>
                  <td key={idx} className="px-10 text-center py-4 font-mono">
                    {row?.Status}
                  </td>
                  <td key={idx} className="px-10 text-center py-4 font-mono">
                    {row?.MessageUsage}
                  </td>
                  {/* );
                })} */}
                </Link>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}

export default QueueCard;
