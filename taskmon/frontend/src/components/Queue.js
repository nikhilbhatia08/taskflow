import React, { useState, useMemo } from "react";
import { useParams } from "react-router-dom";
import Pagination from "./Pagination";

let PageSize = 10;

function Queue() {
  const [data, setData] = useState([1, 2, 3, 4, 5, 6, 7, 8, 9]);
  const [currentPage, setCurrentPage] = useState(1);

  const currentTableData = useMemo(() => {
    const firstPageIndex = (currentPage - 1) * PageSize;
    const lastPageIndex = firstPageIndex + PageSize;
    return data.slice(firstPageIndex, lastPageIndex);
  }, [currentPage]);
  const params = useParams();
  const queuename = params.QueueName;
  return (
    <div className="ml-20">
      <h1 className="text-black">{queuename}</h1>

      <Pagination
        className="pagination-bar"
        currentPage={currentPage}
        totalCount={data.length}
        pageSize={PageSize}
        onPageChange={(page) => setCurrentPage(page)}
      />
    </div>
  );
}

export default Queue;
