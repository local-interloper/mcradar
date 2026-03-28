import { Pagination, Spinner, Table } from "@heroui/react";
import type { Column } from "../types/column";
import React, { useEffect, useRef, useState, type ReactNode } from "react";
import { fetchTableData } from "../utils/fetching-utils";

interface Props<T> {
  endpoint: string;
  columns: Column<T>[];
  filters: any;
}

const PAGE_SIZE = 20;

export function DataTable<T>({ endpoint, columns, filters }: Props<T>) {
  const prevFilters = useRef(filters);

  const abortController = useRef(new AbortController());

  const [page, setPage] = useState(0);
  const [total, setTotal] = useState(0);
  const [tableData, setTableData] = useState<TableData<T> | undefined>(
    undefined,
  );

  const fetchData = async () => {
    const data = await fetchTableData<T>({
      endpoint,
      pagination: {
        first: page * PAGE_SIZE,
        last: page * PAGE_SIZE + PAGE_SIZE,
      },
      filters,
    });

    setTotal(data.total);
    setTableData(data);
  };

  if (prevFilters.current !== filters) {
    prevFilters.current = filters;
    setPage(0);
    setTableData(undefined);
  }

  useEffect(() => {
    fetchData();
  }, [filters, page]);

  const pages = Array.from({ length: Math.ceil(total / PAGE_SIZE) }).map(
    (_, i) => i + 1,
  );

  return tableData ? (
    <Table variant="secondary" className="flex flex-col h-full overflow-hidden">
      <Table.ScrollContainer className="h-full">
        <Table.Content>
          <Table.Header className="sticky top-0 z-10">
            {columns.map(({ label }, i) => (
              <Table.Column key={i} isRowHeader={i === 0}>
                {label}
              </Table.Column>
            ))}
          </Table.Header>
          <Table.Body>
            {tableData &&
              tableData.data.map((row, i) => (
                <Table.Row key={i}>
                  {columns.map(({ property, Formatter }, i) => (
                    <Table.Cell key={i}>
                      {!!Formatter ? (
                        <Formatter>{row[property] as string}</Formatter>
                      ) : (
                        (row[property] as ReactNode)
                      )}
                    </Table.Cell>
                  ))}
                </Table.Row>
              ))}
          </Table.Body>
        </Table.Content>
      </Table.ScrollContainer>
      <Table.Footer className="h-min">
        <Pagination size="sm">
          <Pagination.Content>
            {pages.map((pageNum, i) => (
              <Pagination.Item key={i}>
                <Pagination.Link
                  isActive={page === i}
                  onPress={() => setPage(i)}
                >
                  {pageNum}
                </Pagination.Link>
              </Pagination.Item>
            ))}
          </Pagination.Content>
        </Pagination>
      </Table.Footer>
    </Table>
  ) : (
    <article className="flex justify-center items-center w-full h-full">
      <Spinner size="xl" />
    </article>
  );
}
